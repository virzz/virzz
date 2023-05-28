package jwttool

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var (
	secretFlag = &cli.StringFlag{
		Name:    "secret",
		Aliases: []string{"s"},
		Usage:   "JWT secret",
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
				r, err := JWTPrint(token, c.String("secret"))
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
				&cli.BoolFlag{
					Name:    "none",
					Aliases: []string{"n"},
					Value:   false,
					Usage:   "Set none method and no signature. (Deprecated)",
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
			},
			Action: func(c *cli.Context) (err error) {
				token := c.String("token")
				if token == "" {
					if c.NArg() < 1 {
						return fmt.Errorf("invalid arguments")
					}
					token = c.Args().First()
				}
				r, err := JWTModify(
					token, c.Bool("none"), c.String("secret"),
					c.StringMap("claims"), c.String("method"),
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
				r, err := JWTCreate(
					c.StringMap("claims"), c.String("method"), c.String("secret"),
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
