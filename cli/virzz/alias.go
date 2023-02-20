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
}

var denyAlias = []string{"help", "completion", "alias", "version", "print"}

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
			// alias cmd='prefix cmd'
			res = append(res, fmt.Sprintf("alias %s='%s %s'", c.Name, prefix, c.Name))
		}
	}
	return res
}

var aliasCmd = &cli.Command{
	Name:   "alias",
	Usage:  "alias prefix cmd",
	Hidden: true,
	Action: func(c *cli.Context) error {
		fmt.Println(strings.Join(getCommandAlias(c.App.Name, c.App.Commands), "\n"))
		return nil
	},
}
