package dns

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/miekg/dns"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/services/mariadb"
)

var dnsServer *dns.Server

func Start() {
	if !viper.IsSet("dns") {
		logger.Fatal("Not set mariadb config")
	}
	dnsServer = NewServer()
	go func() {
		if err := dnsServer.ListenAndServe(); err != nil {
			logger.Error("DNS Server Serve Error: ", err)
			dnsServer = nil
		}
	}()
	logger.SuccessF("Running DNS Server Listen on: %s", dnsServer.Addr)
}

func Shutdown(ctx context.Context) {
	if dnsServer != nil {
		if err := dnsServer.ShutdownContext(ctx); err != nil {
			logger.Error("DNS Server Shutdown Error", err)
			return
		}
		logger.Success("DNS Server Shutdown ...")
		return
	}
	logger.Error("DNS Server is not running")
}

var Cmd = &cli.Command{
	Name:    "dns",
	Usage:   "Service DNSLog",
	Aliases: []string{"p"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug-database",
			Usage:   "Enable Database debug mode",
			Aliases: []string{"X"},
		},
	},
	Action: func(c *cli.Context) error {
		err := mariadb.Connect(c.Bool("debug-database"))
		if err != nil {
			return err
		}
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		Start()
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		Shutdown(ctx)
		return nil
	},
}
