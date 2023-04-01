package encrypt

import (
	"path"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

var Cmd = &cli.Command{
	Category: "Crypto",
	Name:     "encrypt",
	Usage:    "Encrypt data by aes and compress by zstd",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "Secret Key",
			Value:   "",
		},
		&cli.StringFlag{
			Name:  "iv",
			Usage: "Secret IV",
			Value: "",
		},
		&cli.BoolFlag{
			Name:  "no-compress",
			Usage: "Encrypt but not compress",
		},
		&cli.BoolFlag{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "Detect whether the file is encrypted",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if c.Bool("check") {
			return Check(c.Args().First())
		}
		key := c.String("key")
		if len(key) == 0 {
			conf := viper.New()
			conf.AddConfigPath(path.Join("$HOME", ".config", "virzz"))
			conf.SetConfigName("encrypt")
			conf.SetConfigType("yaml")
			if err := conf.ReadInConfig(); err != nil {
				logger.Debug(err)
			}
			key = conf.GetString("key")
			if len(key) == 0 {
				logger.Error("key is empty")
				key = utils.RandomStringByLength(32)
				logger.DebugF("key: %s", key)
				conf.Set("key", key)
				if err = conf.WriteConfig(); err != nil {
					logger.Debug(err)
					if err = conf.SafeWriteConfig(); err != nil {
						logger.Debug(err)
					}
				}
				logger.Success("Generate new key and save to config")
			}
		}
		err = Encrypt(c.Args().First(), []byte(key), []byte(c.String("iv")), !c.Bool("no-compress"))
		if err != nil {
			logger.DebugF("%+v", err)
		}
		return err
	},
}
