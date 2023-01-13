package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/services/server/mariadb"
	"github.com/virzz/virzz/services/server/web"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Service Web Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := ServiceWebServer()
		if err != nil {
			return err
		}
		return common.Output(string(r))
	},
}

func webServerStart() {
	if !viper.IsSet("dns") {
		logger.Fatal("Not set mariadb config")
	}
	webServer := web.NewWebServer(debugMode)
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP Server Serve Error: ", err)
		}
	}()
	logger.SuccessF("Run HTTP Server Listen on: %s", webServer.Addr)
	servers = append(servers, webServer)
}

// ServiceWebServer -
func ServiceWebServer() (string, error) {

	err := mariadb.Connect(debugDatabase)
	if err != nil {
		return "", err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	webServerStart()

	<-interrupt
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	serverShutdown(ctx)

	return "", nil
}
