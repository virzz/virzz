package execext

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	var stdout bytes.Buffer
	opts := &RunCommandOptions{
		Command: "ls",
		Dir:     ".",
		Stdout:  &stdout,
		Stderr:  os.Stderr,
	}
	if err := RunCommand(context.Background(), opts); err != nil {
		fmt.Printf(`RunCommand "%s" failed: %s`, opts.Command, err)
		return
	}

	// Trim a single trailing newline from the result to make most command
	// output easier to use in shell commands.
	result := strings.TrimSuffix(strings.TrimSuffix(stdout.String(), "\r\n"), "\n")

	fmt.Println(result)
}
