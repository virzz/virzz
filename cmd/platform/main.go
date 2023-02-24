package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/services/server/dns"
	"github.com/virzz/virzz/services/server/mariadb"
	"github.com/virzz/virzz/services/server/web"
	"github.com/virzz/virzz/utils"
)

const BinName = "virzz-platform"

var (
	Version  string = "latest"
	BuildID  string = "0"
	Revision string = ""
)

func init() {

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/virzz")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Warn(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.Set("modify", true)
			_secret, _ := hash.Sha1Hash([]byte("virzz_jwt_secret"))
			viper.Set("jwt.secret", _secret)
			viper.Set("jwt.expires", 3600)

			viper.Set("mariadb.host", "127.0.0.1")
			viper.Set("mariadb.port", 3306)
			viper.Set("mariadb.charset", "utf8mb4")
			viper.Set("mariadb.name", "virzz_platform")
			viper.Set("mariadb.user", "virzz")
			viper.Set("mariadb.pass", "virzz9999")

			viper.Set("dns.timeout", 5)
			viper.Set("dns.ttl", 600)
			viper.Set("dns.port", 53)
			viper.Set("dns.host", "127.0.0.1")
			viper.Set("dns.domain", "example.com")

			viper.Set("web.host", "127.0.0.1")
			viper.Set("web.port", 8088)

			viper.Set("redis.host", "127.0.0.1")
			viper.Set("redis.port", 6379)
			viper.Set("redis.db", 0)

			if err := viper.SafeWriteConfig(); err != nil {
				logger.Error(err)
				return
			}
			if err := viper.ReadInConfig(); err != nil {
				logger.Fatal(err)
			}
			logger.FatalF("New Config created at '%s'\nYou must modify and remove key 'modify'", viper.ConfigFileUsed())
		} else {
			logger.Error("Config file was found but another error was produced", err)
		}
	}

	if viper.GetBool("modify") {
		logger.Fatal("You must remove the key 'modify'")
	}

	if common.DebugMode {
		logger.SetDebug(true)
	}

}

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Ver: %s (build-%s) revision=%s\n", c.App.Version, BuildID, Revision)
	}
	app := &cli.App{
		Name:                       BinName,
		Authors:                    []any{fmt.Sprintf("%s <%s>", common.Author, common.Email)},
		Usage:                      "The Cyber Swiss Army Knife for platform",
		Version:                    Version,
		Suggest:                    true,
		EnableShellCompletion:      true,
		HideHelpCommand:            true,
		ShellCompletionCommandName: "completion",
		Commands: []*cli.Command{
			mariadb.Cmd,
			dns.Cmd,
			web.Cmd,
		},
	}

	// HideHelpCommand
	utils.HideHelpCommand(app.Commands)

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
	}
}
