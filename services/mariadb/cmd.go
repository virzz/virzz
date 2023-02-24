package mariadb

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
)

var Cmd = &cli.Command{
	Name:    "mariadb",
	Aliases: []string{"db"},
	Usage:   "Service Mariadb",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "Ebable Debug Mode",
			Aliases: []string{"D"},
		},
		&cli.BoolFlag{
			Name:    "migrate",
			Usage:   "Auto migration",
			Aliases: []string{"m"},
		},
		&cli.BoolFlag{
			Name:    "procedure",
			Usage:   "Print create procedure 'auto clean record' sql",
			Aliases: []string{"p"},
		},
		&cli.BoolFlag{
			Name:    "test",
			Usage:   "Test connect database",
			Aliases: []string{"t"},
		},
	},
	Action: func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			return fmt.Errorf("must need at least one flag")
		}
		if c.Bool("procedure") {
			fmt.Println(Procedure())
			return nil
		}
		err := Connect(c.Bool("debug"))
		if err != nil {
			return err
		}
		if c.Bool("test") {
			res, err := ExecSQL(`select 'success';`)
			if err != nil {
				return err
			}
			logger.Success(res)
		} else if c.Bool("migrate") {
			return Migrate()
		}
		return nil
	},
}
