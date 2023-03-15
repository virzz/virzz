package hashpwd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
)

var Cmd = &cli.Command{
	Category: "Crypto",
	Name:     "hashpwd",
	Usage:    "A tool for query password hash offline",
	Commands: []*cli.Command{
		// generate
		&cli.Command{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate password hash form password dict",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "filepath",
					Aliases: []string{"f", "file"},
					Usage:   "password dict file",
				},
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o", "out"},
					Usage:   "Save password hash dict",
					Value:   "hash.dic",
				},
			},
			Action: func(c *cli.Context) (err error) {
				fpath := c.String("filepath")
				fs, err := os.Stat(fpath)
				if err != nil {
					return errors.WithStack(err)
				}
				if fs.IsDir() {
					return errors.Errorf("%s is not a file", fpath)
				}
				return GenerateHashDict(fpath, c.String("output"))
			},
		},
		// lookup
		&cli.Command{
			Name:    "lookup",
			Aliases: []string{"l"},
			Usage:   "Generate password hash form password dict",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "dict",
					Aliases: []string{"f", "file"},
					Usage:   "password hash dict file",
					Value:   "hash.dic",
				},
				&cli.StringFlag{
					Name:    "hash",
					Aliases: []string{"p"},
					Usage:   "password hash",
				},
			},
			Action: func(c *cli.Context) (err error) {
				p := c.String("hash")
				if p == "" {
					return errors.Errorf("plz input password hash")
				}
				fpath := c.String("dict")
				fs, err := os.Stat(fpath)
				if err != nil {
					return errors.WithStack(err)
				}
				if fs.IsDir() {
					return errors.Errorf("%s is not a file", fpath)
				}
				pwd, err := LookupHashDict(fpath, p)
				if err != nil {
					return err
				}
				if pwd == "" {
					return errors.Errorf("Not found hash: %s", p)
				}
				logger.SuccessF("Found password: %s", pwd)
				return nil
			},
		},
	},
}
