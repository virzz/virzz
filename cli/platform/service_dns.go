package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/common"
	vdns "github.com/virzz/virzz/services/server/dns"
	"github.com/virzz/virzz/services/server/mariadb"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Service DNS Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := ServiceDNSServer()
		if err != nil {
			return err
		}
		return common.Output(string(r))
	},
}

func dnsServerStart() {
	if !viper.IsSet("dns") {
		logger.Fatal("Not set mariadb config")
	}

	dnsServer := vdns.NewDNSServer()
	go func() {
		if err := dnsServer.ListenAndServe(); err != nil {
			logger.Error("DNS Server Serve Error: ", err)
		}
	}()
	logger.SuccessF("Running DNS Server Listen on: %s", dnsServer.Addr)
	servers = append(servers, dnsServer)
}

// ServiceDNSServer -
func ServiceDNSServer() (string, error) {

	err := mariadb.Connect(debugDatabase)
	if err != nil {
		return "", err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dnsServerStart()

	<-interrupt
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	serverShutdown(ctx)

	return "", nil
}
