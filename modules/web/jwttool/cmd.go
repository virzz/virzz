package jwttool

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
)

var (
	secretFlag = &cli.StringFlag{
		Name:    "secret",
		Aliases: []string{"s"},
		Usage:   "JWT secret",
	}
	secretFileFlag = &cli.StringFlag{
		Name:    "secret-file",
		Aliases: []string{"sf"},
		Usage:   "JWT secret from file",
	}
	tokenFlag = &cli.StringFlag{
		Name:    "token",
		Aliases: []string{"t"},
		Usage:   "JWT token",
	}
)

var Cmd = &cli.Command{
	Category: "Web",
	Name:     "jwttool",
	Aliases:  []string{"jwt"},
	Usage:    "A jwt tool with Print/Crack/Modify",
	Commands: []*cli.Command{
		// Print
		{
			Category: "JWT",
			Name:     "jwtp",
			Aliases:  []string{"print", "p"},
			Usage:    "Print jwt pretty",
			Flags: []cli.Flag{
				tokenFlag,
				secretFlag,
			},
			Action: func(c *cli.Context) (err error) {
				token := c.String("token")
				if token == "" {
					if c.NArg() < 1 {
						return fmt.Errorf("invalid arguments")
					}
					token = c.Args().First()
				}
				secret := c.String("secret")
				secretFile := c.String("secret-file")
				if secretFile != "" {
					buf, err := os.ReadFile(secretFile)
					if err == nil {
						secret = string(buf)
					} else {
						logger.Warn(err)
					}
				}
				r, err := JWTPrint(token, secret)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// Modify
		{
			Category: "JWT",
			Name:     "jwtm",
			Aliases:  []string{"modify", "m"},
			Usage:    "Modify jwt",
			Flags: []cli.Flag{
				tokenFlag,
				secretFlag,
				secretFileFlag,
				&cli.BoolFlag{
					Name:    "none",
					Aliases: []string{"n"},
					Value:   false,
					Usage:   "Set none method and no signature. (Deprecated)",
				},
				&cli.BoolFlag{
					Name:  "print",
					Value: false,
					Usage: "JWTPrint result token",
				},
				&cli.StringFlag{
					Name:    "method",
					Aliases: []string{"m"},
					Usage:   "Set new method: <HS256|HS384|HS512>",
				},
				&cli.StringMapFlag{
					Name:    "claims",
					Aliases: []string{"c"},
					Usage:   "modify or add claims",
				},
				&cli.StringMapFlag{
					Name:    "headers",
					Aliases: []string{"p", "H"},
					Usage:   "modify or add headers",
				},
			},
			Action: func(c *cli.Context) (err error) {
				token := c.String("token")
				if token == "" {
					if c.NArg() < 1 {
						return fmt.Errorf("invalid arguments")
					}
					token = c.Args().First()
				}
				secret := c.String("secret")
				secretFile := c.String("secret-file")
				if secretFile != "" {
					buf, err := os.ReadFile(secretFile)
					if err == nil {
						secret = string(buf)
					} else {
						logger.Warn(err)
					}
				}
				r, err := JWTModify(
					token, c.Bool("none"), secret,
					c.StringMap("headers"), c.StringMap("claims"), c.String("method"),
					c.Bool("print"),
				)
				if err != nil {
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// Crack
		{
			Category: "JWT",
			Name:     "jwtc",
			Aliases:  []string{"crack", "c"},
			Usage:    "Crack jwt",
			Flags: []cli.Flag{
				tokenFlag,
				&cli.StringFlag{
					Name:    "alphabet",
					Aliases: []string{"a"},
					Value:   defaultAlphabet,
					Usage:   "the alphabet for the brute",
				},
				&cli.StringFlag{
					Name:    "prefix",
					Aliases: []string{"p"},
					Usage:   "prefixed to the secret",
				},
				&cli.StringFlag{
					Name:    "suffix",
					Aliases: []string{"s"},
					Usage:   "suffixed to the secret",
				},
				&cli.IntFlag{
					Name:    "minlen",
					Aliases: []string{"m"},
					Usage:   "The min length secret",
				},
				&cli.IntFlag{
					Name:    "maxlen",
					Aliases: []string{"l"},
					Usage:   "The max length secret",
				},
			},
			Action: func(c *cli.Context) (err error) {
				token := c.String("token")
				if token == "" {
					if c.NArg() < 1 {
						return fmt.Errorf("invalid arguments")
					}
					token = c.Args().First()
				}
				r, err := JWTCrack(
					token, c.Int("minlen"), c.Int("maxlen"),
					c.String("alphabet"), c.String("prefix"), c.String("suffix"),
				)
				if err != nil {
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// Create/Generate
		{
			Category: "JWT",
			Name:     "jwtg",
			Aliases:  []string{"generate", "create", "n"},
			Usage:    "Create/Generate jwt",
			Flags: []cli.Flag{
				secretFlag,
				secretFileFlag,
				&cli.StringFlag{
					Name:    "method",
					Aliases: []string{"m"},
					Usage:   "Set new method: <HS256|HS384|HS512>",
				},
				&cli.StringMapFlag{
					Name:    "claims",
					Aliases: []string{"c"},
					Usage:   "modify or add  claims/payload",
				},
			},
			Action: func(c *cli.Context) (err error) {
				secret := c.String("secret")
				secretFile := c.String("secret-file")
				if secretFile != "" {
					buf, err := os.ReadFile(secretFile)
					if err == nil {
						secret = string(buf)
					} else {
						logger.Warn(err)
					}
				}
				r, err := JWTCreate(
					c.StringMap("claims"), c.String("method"), secret,
				)
				if err != nil {
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},
	},
}
