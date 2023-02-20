package hash

import "github.com/spf13/cobra"

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		hashCmd,
		hmacCmd,
		// gmsmCmd,
		// bcryptCmd,
	}
}
