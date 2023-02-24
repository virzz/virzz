package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/services/mariadb"
)

var httpServer *http.Server

func Start(debug bool) {
	if !viper.IsSet("web") {
		logger.Fatal("Not set web config")
	}
	httpServer = NewServer(debug)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error("HTTP Server Serve Error: ", err)
			httpServer = nil
		}
	}()
	logger.SuccessF("Running HTTP Server Listen on: %s", httpServer.Addr)
}

func Shutdown(ctx context.Context) {
	if httpServer != nil {
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error("HTTP Server Shutdown Error", err)
			return
		}
		logger.Success("HTTP Server Shutdown ...")
		return
	}
	logger.Error("HTTP Server is not running")
}

var Cmd = &cli.Command{
	Name:    "web",
	Usage:   "Service WebLog",
	Aliases: []string{"w"},
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
	},
	Action: func(c *cli.Context) error {

		err := mariadb.Connect(c.Bool("debug-database"))
		if err != nil {
			return err
		}

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		Start(c.Bool("debug-mode"))

		<-interrupt

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		Shutdown(ctx)

		return nil
	},
}
