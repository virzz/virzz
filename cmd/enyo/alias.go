package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

var allowAlias = []string{
	"jwttool",
	"domain",
	"basex",
	"basic",
	"classical",
	"hash",
	"netool",
	"parser",
	"qrcode",
}

var denyAlias = []string{
	"help",
	"completion",
	"alias",
	"version",
}

func sliceContains[T comparable](inputSlice []T, element T) bool {
	for _, inputValue := range inputSlice {
		if inputValue == element {
			return true
		}
	}
	return false
}

func getCommandAlias(prefix string, cmds []*cli.Command) []string {
	var res = make([]string, 0)
	for _, c := range cmds {
		if len(c.Commands) > 0 && sliceContains(allowAlias, c.Name) {
			res = append(res, getCommandAlias(fmt.Sprintf("%s %s", prefix, c.Name), c.Commands)...)
		} else if sliceContains(denyAlias, c.Name) {
			continue
		} else {
			res = append(res, fmt.Sprintf("which %s >/dev/null || alias %s='%s %s'", c.Name, c.Name, prefix, c.Name))
		}
	}
	return res
}

var aliasCmd = &cli.Command{
	Name:   "alias",
	Usage:  "alias prefix cmd, use -r to show .*shrc script",
	Hidden: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "rc",
			Aliases: []string{"r"},
			Usage:   "Generate alias for shell config",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("rc") {
			fmt.Printf(`
if which %s > /dev/null ; then
    source <(%s alias)
    PROG=%s source <(%s completion zsh)
fi
`, BinName, BinName, BinName, BinName)
			return nil
		}
		fmt.Println(strings.Join(getCommandAlias(c.App.Name, c.App.Commands), "\n"))
		return nil
	},
}
