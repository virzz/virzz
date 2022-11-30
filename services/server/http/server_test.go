package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/logger"
	"github.com/mozhu1024/virzz/services/server/mariadb"
)

func TestConfig(t *testing.T) {
	data := common.TemplateConfig()
	fmt.Println(string(data))
}

func TestServer(t *testing.T) {
	err := common.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	err = mariadb.Connect()
	if err != nil {
		t.Fatal(err)
	}
	httpConfig := &Config{
		Host:  "127.0.0.1",
		Port:  9999,
		Debug: true,
	}
	// Run HTTP Server
	httpServer := NewWebServer(httpConfig)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server listen: ", err)
		}
	}()
	logger.Debug("HTTP Server Running on ", httpServer.Addr)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error("Shutdown Error", err)
		}
	}()

	time.Sleep(5 * time.Second)

	resp, err := http.Get("http://127.0.0.1:9999/ping")
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
