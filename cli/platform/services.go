package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/services/server/mariadb"
)

var (
	servers    []interface{}
	dnsService bool
	webService bool
)

var platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Run Service Platform",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := ServicePlatform()
		if err != nil {
			return err
		}
		return common.Output(string(r))
	},
}

func init() {
	platformCmd.Flags().BoolVarP(&dnsService, "dns", "d", false, "Run DNS Service")
	platformCmd.Flags().BoolVarP(&webService, "web", "w", false, "Run HTTP Service")
}

func serverShutdown(ctx context.Context) {
	for _, s := range servers {
		switch _s := s.(type) {
		case *http.Server:
			if err := _s.Shutdown(ctx); err != nil {
				logger.Error("Web Server Shutdown Error", err)
			}
			logger.Success("Web Server Shutdown ...")
		case *dns.Server:
			if err := _s.ShutdownContext(ctx); err != nil {
				logger.Error("DNS Server Shutdown Error", err)
			}
			logger.Success("DNS Server Shutdown ...")
		default:
			continue
		}
	}
}

// ServicePlatform -
func ServicePlatform() (string, error) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	err := mariadb.Connect(debugDatabase)
	if err != nil {
		return "", err
	}

	if dnsService {
		dnsServerStart()
	}
	if webService {
		webServerStart()
	}

	sig := <-interrupt
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	serverShutdown(ctx)

	if sig == os.Interrupt {
		logger.Info("Platform was interruped by system signal")
	} else {
		logger.Info("Platform was killed")
	}

	return "", nil
}
