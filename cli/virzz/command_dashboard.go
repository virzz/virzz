//go:build dashboard
// +build dashboard

package main

import (
	"github.com/virzz/virzz/modules/dashboard"
)

func init() {
	// Add SubCommands
	commands = append(commands, dashboard.Cmd)
}
