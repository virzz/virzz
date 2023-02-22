package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils"
)

const (
	TARGET_DIR  = "./build"
	SOURCE_DIR  = "./cli"
	PUBLIC_DIR  = "./cli/public"
	RELEASE_DIR = "./release"
	MULTI_COUNT = 6
)

var (
	OSS    = []string{"linux", "windows", "darwin"}
	ARCHES = []string{"amd64", "arm64"}
)

func main() {
	app := &cli.App{
		Name:      "compile",
		Authors:   []any{fmt.Sprintf("%s <%s>", common.Author, common.Email)},
		Usage:     "Helper for building Virzz",
		UsageText: `go run ./internal/_compile [flags] name...`,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "clean", Aliases: []string{"C"}, Usage: "Clean build files"},
			&cli.BoolFlag{Name: "release", Aliases: []string{"R"}, Usage: "Build release version"},
			&cli.BoolFlag{Name: "archive", Aliases: []string{"A"}, Usage: "Archive release packages"},
			&cli.BoolFlag{Name: "together", Aliases: []string{"t"}, Usage: "Archive all in one"},
			&cli.BoolFlag{Name: "multi", Aliases: []string{"M"}, Usage: "Compile multi-platform binaries"},
			&cli.BoolFlag{Name: "force", Aliases: []string{"F"}, Usage: "Force to Compile"},
			&cli.StringFlag{Name: "version", Aliases: []string{"V"}, Usage: "Custom Build version"},
			&cli.StringFlag{Name: "output", Aliases: []string{"O"}, Usage: "Custom output path", Value: TARGET_DIR},
			&cli.StringFlag{Name: "builder", Aliases: []string{"B"}, Usage: "Replace `go` builder", Value: "go"},
			&cli.StringSliceFlag{Name: "go-tags", Aliases: []string{"T"}, Usage: "Append go build tags"},
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "Print verbose information"},
			&cli.BoolFlag{Name: "verbosex", Aliases: []string{"vv"}, Usage: "Print verbose information"},
		},
		Action: func(c *cli.Context) error {

			// Clean Build Directory
			if c.Bool("clean") {
				if err := os.RemoveAll(TARGET_DIR); err != nil {
					return err
				}
				if err := os.Mkdir(TARGET_DIR, 0775); err != nil {
					return err
				}
				logger.Success("Cleaned build directory")
			}

			if c.NArg() == 0 {
				return fmt.Errorf("plz input args:{project...}")
			}

			if c.Bool("verbosex") {
				verbose = VerboseX
			} else if c.Bool("verbose") {
				verbose = VerboseV
			} else {
				verbose = 0
			}

			force = c.Bool("force")
			release = c.Bool("release")
			goVersion = c.String("version")
			goOutput = c.String("output")
			goBuilder = c.String("builder")
			goTags = append(goTags, c.StringSlice("go-tags")...)

			projs := c.Args().Slice()
			// Compile Virzz Projects
			sourceNames := make(map[string]string)
			publicNames := getPublicProjects()
			// all public
			if utils.SliceContains(projs, "public") {
				for _, name := range publicNames {
					sourceNames[name] = fmt.Sprintf("%s/%s", PUBLIC_DIR, name)
				}
			}
			// inner public
			for _, name := range utils.Intersection(publicNames, projs) {
				sourceNames[name] = fmt.Sprintf("%s/%s", PUBLIC_DIR, name)
			}
			// specific
			for _, name := range utils.Intersection(getSpecialProjects(), projs) {
				sourceNames[name] = fmt.Sprintf("%s/%s", SOURCE_DIR, name)
			}

			// Compile
			for name, source := range sourceNames {
				buildID, err := BuildID.Inc(name)
				if err != nil {
					logger.Error(err)
				}

				if c.Bool("multi") {
					logger.InfoF("Start compiling %s for all platforms", name)
					if len(multiCompile(name, source, buildID)) == 6 {
						logger.Success("All platforms compiled successfully")
						if c.Bool("archive") {
							if c.Bool("together") {
								logger.Info("Start archiving togher")
							} else {
								logger.Info("Start archiving")
							}
							archiveTargets(name, c.Bool("together"))
						}
					} else {
						logger.Error("Lost some platform binaries, try compile again")
					}

				} else {
					logger.InfoF("Start compiling %s", name)
					if err := compile(name, source, name, buildID); err != nil {
						logger.ErrorF("Compile %s fail: %v", name, err)
						continue
					}
					logger.SuccessF("Compiled [%s] successfully", name)
				}
			}

			return nil
		},
	}

	// HideHelpCommand
	utils.HideHelpCommand(app.Commands)

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
	}
}
