package cmd

import (
	"github.com/spf13/cobra"
	cm "github.com/virink/virzz/common"
	"github.com/virink/virzz/misc/collision/crc32"
)

func init() {
	var (
		table  string
		length int
	)

	// collisionCmd
	var collisionCmd = &cobra.Command{
		Use:   "collision",
		Short: "Collision",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// zipCrc32Cmd
	var zipCrc32Cmd = &cobra.Command{
		Use:   "zipcrc [filename]",
		Short: "Zip CRC32",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := crc32.ZipCRC32(args[0], table, length)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	zipCrc32Cmd.Flags().IntVarP(&length, "length", "l", 4, "Uncompressed Size Limit")
	zipCrc32Cmd.Flags().StringVarP(&table, "table", "t", "abcdefghijklmnopqrstuvwxyz", "")
	collisionCmd.AddCommand(zipCrc32Cmd)
	rootCmd.AddCommand(collisionCmd)
}
