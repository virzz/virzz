package ghext

import (
	"fmt"
	"os"
	"path"
	"strings"

	ghConfig "github.com/cli/go-gh/pkg/config"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

func _prompt(isHideEmoji bool) (err error) {
	err = prompt()
	if err != nil {
		return
	}
	res, err := CommitTemplate(
		int(cmsg.Type), cmsg.Scope, cmsg.Subject, cmsg.Body, cmsg.Footer, isHideEmoji)
	if err != nil {
		return err
	}
	_, err = fmt.Print(res)
	return err
}

func init() {
	// commit - Generate Commit Message
	commitCmd := &cli.Command{
		Name:    "commit",
		Usage:   "Generate Commit Message",
		Aliases: []string{"gcmt"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "scope",
				Usage: "Scope",
				Value: "-",
			},
			&cli.StringFlag{
				Name:  "subject",
				Usage: "subject",
			},
			&cli.StringFlag{
				Name:  "body",
				Usage: "body",
			},
			&cli.StringFlag{
				Name:  "footer",
				Usage: "footer",
			},
			&cli.BoolFlag{
				Name:  "hide-emoji",
				Usage: "Hide Emoji",
			},
			&cli.BoolFlag{
				Name:  "prompt",
				Usage: "Prompt",
			},
		},
		Action: func(c *cli.Context) (err error) {
			var res = ""
			if c.Bool("prompt") {
				return _prompt(c.Bool("hide-emoji"))
			}

			typ := 0
			for i := 1; i < len(_MsgType_index); i++ {
				if c.Bool(strings.ToLower(MsgType(i).String())) {
					typ = i
					break
				}
			}
			// typ, scope, subject, body, footer
			res, err = CommitTemplate(
				typ,
				c.String("scope"),
				c.String("subject"),
				c.String("body"),
				c.String("footer"),
				c.Bool("hide-emoji"))
			if err != nil {
				return err
			}
			_, err = fmt.Print(res)
			return
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "prompt",
				Usage:   "Prompt",
				Aliases: []string{"p"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "hide-emoji",
						Usage: "Hide Emoji",
					},
				},
				Action: func(c *cli.Context) (err error) {
					return _prompt(c.Bool("hide-emoji"))
				},
			},
		},
	}

	for i := 1; i < len(_MsgType_index); i++ {
		commitCmd.Flags = append(commitCmd.Flags, &cli.BoolFlag{
			Name:  strings.ToLower(MsgType(i).String()),
			Usage: fmt.Sprintf("Commit Type %s", MsgType(i)),
		})
		commitTypeItems = append(commitTypeItems, MsgType(i).String())
	}
	Cmd.Commands = append(Cmd.Commands, commitCmd)
}

var Cmd = &cli.Command{
	Category: "GitHub",
	Name:     "ghext", // For GitHub command-line tool
	Aliases:  []string{"gh-mozhu"},
	Usage:    "A little toolkit using GitHub API",
	Commands: []*cli.Command{
		// install - Install this
		&cli.Command{
			Name:  "install",
			Usage: "Install this",
			Action: func(c *cli.Context) (err error) {
				binName := path.Base(os.Args[0])
				if !strings.HasPrefix(binName, "gh-") {
					err = fmt.Errorf("extension name [%s] is error", binName)
					return
				}
				// Local extensions
				// $extdir/repo-name/bin-name repo-name == bin-name
				newPath := path.Join(ghConfig.DataDir(), "extensions", binName, binName)
				logger.Debug("gh extensions dir: ", newPath)
				err = utils.CopyFile(newPath, os.Args[0])
				if err != nil {
					return
				}
				logger.SuccessF("Installed: %s", newPath)
				return
			},
		},
		// orgs - List organizations for the authenticated user
		&cli.Command{
			Name:  "orgs",
			Usage: "List organizations for the authenticated user",
			Action: func(c *cli.Context) (err error) {
				res, err := ListUserOrganizations()
				if err != nil {
					return
				}
				fmt.Println(res)
				return
			},
		},
		// transfer - Transfer a repository
		&cli.Command{
			Name:  "transfer",
			Usage: "Transfer a repository",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "owner",
					Aliases:  []string{"o"},
					Usage:    "The account owner of the repository",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "repo",
					Aliases:  []string{"r"},
					Usage:    "The name of the repository",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "new-owner",
					Aliases:  []string{"new", "n"},
					Usage:    "The username/organization whitch will be transferred to.",
					Required: true,
				},
			},
			Action: func(c *cli.Context) (err error) {
				res, err := TransferRepository(c.String("new-owner"), c.String("owner"), c.String("repo"))
				if err != nil {
					return
				}
				logger.Success(res)
				return
			},
		},
	},
}