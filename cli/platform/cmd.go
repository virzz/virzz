package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/services/server/dns"
	"github.com/virzz/virzz/services/server/mariadb"
	"github.com/virzz/virzz/services/server/web"
)

var Cmd = &cli.Command{
	Name:    "platform",
	Usage:   "Service Platform",
	Aliases: []string{"p"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug-mode",
			Usage:   "Enable debug mode",
			Aliases: []string{"D"},
		},
		&cli.BoolFlag{
			Name:    "debug-database",
			Usage:   "Enable Database debug mode",
			Aliases: []string{"X"},
		},
		&cli.BoolFlag{
			Name:    "no-dns",
			Usage:   "Disable DNS service",
			Aliases: []string{"nodns"},
		},
		&cli.BoolFlag{
			Name:    "no-web",
			Usage:   "Disable HTTP service",
			Aliases: []string{"noweb"},
		},
	},
	Action: func(c *cli.Context) error {

		err := mariadb.Connect(c.Bool("debug-database"))
		if err != nil {
			return err
		}

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		if !c.Bool("no-dns") {
			dns.Start()
		}
		if !c.Bool("no-web") {
			web.Start(c.Bool("debug-mode"))
		}

		sig := <-interrupt

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if !c.Bool("no-dns") {
			dns.Shutdown(ctx)
		}
		if !c.Bool("no-web") {
			web.Shutdown(ctx)
		}

		if sig == os.Interrupt {
			logger.Success("Platform was interruped by system signal")
		} else {
			logger.Error("Platform was killed")
		}

		return nil
	},
}
