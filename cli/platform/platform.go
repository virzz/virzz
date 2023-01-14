package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/hash"
)

var (
	AppName        = "VirzzPlatform"
	BinName        = "platform"
	Version string = "dev"
	BuildID string = "0"
)

var versionCmd = common.VersionCommand(AppName, Version, BuildID)

var rootCmd = &cobra.Command{
	Use:           BinName,
	Short:         "The Cyber Swiss Army Knife for platform",
	SilenceErrors: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var (
	debugMode     bool
	debugDatabase bool
	// cacheRedis    bool
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

			viper.Set("jwt.secret", hash.EMd5Hash([]byte("virzz_jwt_secret")))
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

	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "D", false, "Set Debug Mode")
	rootCmd.PersistentFlags().BoolVarP(&debugDatabase, "database", "X", false, "Set Database Debug Mode")
	// rootCmd.PersistentFlags().BoolVarP(&cacheRedis, "cache-redis", "C", false, "Use Redis Cache")

	rootCmd.AddCommand(
		versionCmd,
		platformCmd,

		mariadbCmd,
		dnsCmd,
		webCmd,
	)
}

func main() {
	if common.DebugMode || debugMode {
		logger.SetDebug(true)
	}
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
