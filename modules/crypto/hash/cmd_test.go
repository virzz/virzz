package hash

import (
	"testing"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
)

const (
	plainText = "Let life be beautiful like summer flowers And Death like autumn leaves."
	hamcKey   = "virzz"
)

var Commands = [][]string{
	[]string{"md", "-t", "2"},
	[]string{"md", "-t", "2"},
	[]string{"md", "-t", "4"},
	[]string{"md", "-t", "4"},
	[]string{"md", "-t", "5"},
	[]string{"md", "-t", "5"},

	[]string{"sha1", "", ""},
	[]string{"sha1", "", ""},

	[]string{"sha2", "-t", "224"},
	[]string{"sha2", "-t", "224"},
	[]string{"sha2", "-t", "256"},
	[]string{"sha2", "-t", "256"},
	[]string{"sha2", "-t", "384"},
	[]string{"sha2", "-t", "384"},
	[]string{"sha2", "-t", "512"},
	[]string{"sha2", "-t", "512"},
	[]string{"sha2", "-t", "512224"},
	[]string{"sha2", "-t", "512224"},
	[]string{"sha2", "-t", "512256"},
	[]string{"sha2", "-t", "512256"},

	[]string{"sha3", "-t", "224"},
	[]string{"sha3", "-t", "224"},
	[]string{"sha3", "-t", "256"},
	[]string{"sha3", "-t", "256"},
	[]string{"sha3", "-t", "384"},
	[]string{"sha3", "-t", "384"},
	[]string{"sha3", "-t", "512"},
	[]string{"sha3", "-t", "512"},
}

func TestCmd(t *testing.T) {
	app := &cli.App{Name: "test"}
	base := []string{"test", "hash"}
	app.Commands = append(app.Commands, Cmd)
	for _, cmd := range Commands {
		logger.InfoF("%s %s", cmd[0], cmd[2])
		app.Run(append(base, append(cmd, plainText)...))
		logger.InfoF("%s %s --hmac %s \"%s\"", cmd[0], cmd[2], hamcKey, plainText)
		app.Run(append(base, append(cmd, "-s", hamcKey, plainText)...))
	}

}
