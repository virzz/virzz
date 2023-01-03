package classical

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

var (
	decode   bool = false
	sep, key string
)

var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "Caesar Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, _ := Caesar(s)
		return common.Output(r)
	},
}

var rot13Cmd = &cobra.Command{
	Use:   "rot13",
	Short: "Rot13 By Caesar Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, _ := Rot13(s)
		return common.Output(r)
	},
}

var morseCmd = &cobra.Command{
	Use:   "morse",
	Short: "Morse Code 摩尔斯电码",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Morse(s, decode, sep)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var atbashCmd = &cobra.Command{
	Use:   "atbash",
	Short: "Atbash 埃特巴什码",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Atbash(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var peigenCmd = &cobra.Command{
	Use:   "peigen",
	Short: "Peigen 培根密码",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Peigen(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var vigenereCmd = &cobra.Command{
	Use:   "vigenere",
	Short: "Vigenere 维吉利亚密码",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Vigenere(s, key, decode)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var classicalCmd = &cobra.Command{
	Use:   "classical",
	Short: "Some classical cryptography",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	morseCmd.Flags().BoolVarP(&decode, "decode", "d", false, "Decode")
	morseCmd.Flags().StringVarP(&sep, "sep", "s", "/", "Delimiter")
	vigenereCmd.Flags().StringVarP(&key, "key", "k", "MOZHU", "Vigenere Key")

	classicalCmd.AddCommand(
		caesarCmd,
		rot13Cmd,
		morseCmd,
		atbashCmd,
		peigenCmd,
		vigenereCmd,
	)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		classicalCmd,
	}
}
