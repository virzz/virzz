package web

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/services/server/mariadb"
)

func TestServer(t *testing.T) {

	viper.Set("mariadb", map[string]interface{}{
		"host":    "127.0.0.1",
		"port":    3306,
		"name":    "virzz_platform",
		"user":    "virzz",
		"pass":    "virzz9999",
		"charset": "utf8mb4",
	})
	viper.Set("web", map[string]interface{}{
		"host": "127.0.0.1",
		"port": 8088,
	})

	err := mariadb.Connect()
	if err != nil {
		t.Fatal(err)
	}
	// Run HTTP Server
	webServer := NewWebServer(true)
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Error("HTTP server listen: ", err)
		}
	}()
	logger.Success("HTTP Server Running on ", webServer.Addr)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			t.Error("Shutdown Error", err)
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
