package dashboard

import (
	"github.com/urfave/cli/v3"
)

var Cmd = &cli.Command{
	Category: "Dashboard",
	Name:     "dashboard",
	Usage:    "Dashboard for platform",
	Action: func(c *cli.Context) (err error) {
		DashboardDemo()
		return nil
	},
}
