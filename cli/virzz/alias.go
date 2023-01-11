package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

var allowAlias = []string{
	"virzz",

	"basex",
	"basic",
	"classical",
	"hash",
	// "hmac",

	"netool",
}

func sliceContains[T comparable](inputSlice []T, element T) bool {
	for _, inputValue := range inputSlice {
		if inputValue == element {
			return true
		}
	}
	return false
}

func getCommandAlias(prefix string, cmd *cobra.Command) []string {
	var res = make([]string, 0)
	for _, c := range cmd.Commands() {
		if c.HasSubCommands() && sliceContains(allowAlias, c.Name()) {
			res = append(res, getCommandAlias(fmt.Sprintf("%s %s", prefix, c.Name()), c)...)
		} else if sliceContains([]string{"help", "completion", "alias", "version"}, c.Name()) {
			continue
		} else {
			// alias cmd='prefix cmd'
			res = append(res, fmt.Sprintf("alias %s='%s %s'", c.Name(), prefix, c.Name()))
		}
	}
	return res
}

func aliasCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "alias",
		Short: "Print the version",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd = cmd.Root()
			res := getCommandAlias(cmd.CommandPath(), cmd)
			return common.Output(strings.Join(res, "\n"))
		},
	}
}
