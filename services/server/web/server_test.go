package web

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/services/server/mariadb"
)

func TestConfig(t *testing.T) {
	data, err := common.TemplateConfig()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestServer(t *testing.T) {
	err := mariadb.Connect()
	if err != nil {
		t.Fatal(err)
	}
	// Run HTTP Server
	webServer := NewWebServer(false)
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server listen: ", err)
		}
	}()
	logger.Debug("HTTP Server Running on ", webServer.Addr)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			logger.Error("Shutdown Error", err)
		}
	}()

	time.Sleep(5 * time.Second)

	resp, err := http.Get("http://127.0.0.1:8088/ping")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
