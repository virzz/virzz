package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/logger"
	dnsserver "github.com/mozhu1024/virzz/services/server/dns"
	httpserver "github.com/mozhu1024/virzz/services/server/http"
	"github.com/mozhu1024/virzz/services/server/mariadb"
	"github.com/mozhu1024/virzz/services/server/models"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print Example Config Template",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := common.TemplateConfig()
		if err != nil {
			return err
		}
		return common.Output(string(r))
	},
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run Service Daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := ServiceDaemon()
		if err != nil {
			return err
		}
		return common.Output(string(r))
	},
}

// ServiceDaemon -
func ServiceDaemon() (string, error) {
	// Load Config
	logger.Debug("Load Config ...")
	err := common.LoadConfig()
	if err != nil {
		return "", err
	}

	// Connect Database
	logger.Debug("Database mariadb Connect ...")
	err = mariadb.Connect()
	if err != nil {
		return "", err
	}

	models.InitMariadb()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Run DNS Server
	logger.Debug("Run DNS Server ...")
	dnsServer := dnsserver.NewDNSServer()
	go func() {
		if err := dnsServer.ListenAndServe(); err != nil {
			logger.Error("DNS Server Serve Error: ", err)
		}
	}()
	logger.InfoF("Run DNS Server Listen on: %s", dnsServer.Addr)

	// Run HTTP Server
	logger.Debug("Run HTTP Server ...")
	httpServer := httpserver.NewWebServer()
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP Server Serve Error: ", err)
		}
	}()
	logger.InfoF("Run HTTP Server Listen on: %s", httpServer.Addr)

	sig := <-interrupt

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Shutdown DNS Server
	if err := dnsServer.ShutdownContext(ctx); err != nil {
		logger.Error("DNS Server Shutdown Error", err)
	}
	logger.Info("DNS Server Shutdown ...")

	// Shutdown HTTP Server
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("HTTP Server Shutdown Error", err)
	}
	logger.Info("HTTP Server Shutdown ...")

	if sig == os.Interrupt {
		return "Daemon was interruped by system signal", nil
	}
	return "Daemon was killed", nil
}
