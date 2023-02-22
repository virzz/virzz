package nettool

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var Cmd = &cli.Command{
	Category: "Misc",
	Name:     "netool",
	Usage:    "Net utils for IP/Port",
	Commands: []*cli.Command{
		// ip2oct
		&cli.Command{
			Category: "IP",
			Name:     "ip2oct",
			Aliases:  []string{"oct"},
			Usage:    "IPv4 -> Oct",
			Action: func(c *cli.Context) (err error) {
				ip := c.Args().First()
				if err := utils.ValidArg(ip, "ip"); err != nil {
					return err
				}
				r, err := IP2Oct(ip)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// ip2dec
		&cli.Command{
			Category: "IP",
			Name:     "ip2dec",
			Aliases:  []string{"dec"},
			Usage:    "IPv4 -> Dec",
			Action: func(c *cli.Context) (err error) {
				ip := c.Args().First()
				if err := utils.ValidArg(ip, "ip"); err != nil {
					return err
				}
				r, err := IP2Dec(ip)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// ip2hex
		&cli.Command{
			Category: "IP",
			Name:     "ip2hex",
			Aliases:  []string{"hex"},
			Usage:    "IPv4 -> Hex",
			Action: func(c *cli.Context) (err error) {
				ip := c.Args().First()
				if err := utils.ValidArg(ip, "ip"); err != nil {
					return err
				}
				r, err := IP2Hex(ip)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// ip2all
		&cli.Command{
			Category: "IP",
			Name:     "ip2all",
			Aliases:  []string{"all"},
			Usage:    "IPv4 -> All {Hex|Dec|Oct|...}",
			Action: func(c *cli.Context) (err error) {
				ip := c.Args().First()
				if err := utils.ValidArg(ip, "ip"); err != nil {
					return err
				}
				r, err := IP2All(ip)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// hex2ip
		&cli.Command{
			Category: "IP",
			Name:     "hex2ip",
			Usage:    "Hex -> IPv4",
			Action: func(c *cli.Context) (err error) {
				if c.NArg() == 0 {
					return fmt.Errorf("invlid hex ip")
				}
				r, err := Hex2IP(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// oct2ip
		&cli.Command{
			Category: "IP",
			Name:     "oct2ip",
			Usage:    "Oct -> IPv4",
			Action: func(c *cli.Context) (err error) {
				if c.NArg() == 0 {
					return fmt.Errorf("invlid oct ip")
				}
				r, err := Oct2IP(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// dec2ip
		&cli.Command{
			Category: "IP",
			Name:     "dec2ip",
			Usage:    "Dec -> IPv4",
			Action: func(c *cli.Context) (err error) {
				if c.NArg() == 0 {
					return fmt.Errorf("invlid dec ip")
				}
				r, err := Dec2IP(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// any2ip
		&cli.Command{
			Category: "IP",
			Name:     "dec2ip",
			Usage:    "Any {Hex|Oct|Dec|...} -> IPv4",
			Action: func(c *cli.Context) (err error) {
				if c.NArg() == 0 {
					return fmt.Errorf("invlid special ip")
				}
				r, err := Dec2IP(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// mac2dec
		&cli.Command{
			Category: "MAC",
			Name:     "mac2dec",
			Usage:    "MAC Address -> Dec",
			Action: func(c *cli.Context) (err error) {
				mac := c.Args().First()
				if err := utils.ValidArg(mac, "mac"); err != nil {
					return err
				}
				r, err := Mac2Dec(mac)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// dec2mac
		&cli.Command{
			Category: "MAC",
			Name:     "dec2mac",
			Usage:    "Dec -> MAC Address",
			Action: func(c *cli.Context) (err error) {
				if c.NArg() == 0 {
					return fmt.Errorf("invlid dec mac")
				}
				r, err := Dec2MAC(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
	},
}
