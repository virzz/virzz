package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	_ "github.com/virzz/virzz/common"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/execext"
)

const (
	PACKAGE     = "github.com/virzz/virzz"
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

func getPublicProjects() (names []string) {
	fs, err := os.ReadDir(PUBLIC_DIR)
	if err != nil {
		logger.Error(err)
		return
	}
	names = make([]string, len(fs))
	for i, f := range fs {
		if f.IsDir() {
			names[i] = f.Name()
		}
	}
	return
}

func getSpecialProjects() (names []string) {
	fs, err := os.ReadDir(SOURCE_DIR)
	if err != nil {
		logger.Error(err)
		return
	}
	names = make([]string, len(fs))
	for i, f := range fs {
		if f.IsDir() && f.Name() != "public" && f.Name() != "_compile" {
			names[i] = f.Name()
		}
	}
	return
}

func getVersion(prefix string) string {
	var stdout bytes.Buffer
	opts := &execext.RunCommandOptions{
		Command: fmt.Sprintf("git tag | grep %s | tail -n 1", prefix),
		Dir:     ".",
		Stdout:  &stdout,
		Stderr:  execext.DevNull{},
	}
	if err := execext.RunCommand(context.Background(), opts); err != nil {
		return "error"
	}
	return strings.TrimSuffix(strings.TrimSuffix(stdout.String(), "\r\n"), "\n")
}

func compile(name, source, target string, buildID int) error {

	flags := make(map[string]string)
	if force {
		flags["-a"] = ""
	}
	// -vv print the go build -x commands
	if verbose == 1 {
		flags["-v"] = ""
	} else if verbose == 2 {
		flags["-x"] = ""
	}

	if release {
		version := ""
		if name == "virzz" || name == "platform" {
			version = getVersion("p")
		} else if name != "public" {
			version = getVersion("v")
		}
		flags["-trimpath"] = ""
		flags["-tags"] = "release"
		flags["-ldflags"] = fmt.Sprintf("-s -w -X %s/common.Mode=prod -X main.BuildID=%d -X main.Version=%s", PACKAGE, buildID, version)
	} else {
		flags["-tags"] = "debug"
		flags["-ldflags"] = fmt.Sprintf("-X %s/common.Mode=dev -X main.BuildID=%d", PACKAGE, buildID)
	}

	var flagString bytes.Buffer
	for k, v := range flags {
		flagString.WriteByte(' ')
		if v == "" {
			flagString.WriteString(k)
		} else {
			flagString.WriteString(fmt.Sprintf("%s '%s'", k, v))
		}
	}

	// var stdout bytes.Buffer
	var stderr bytes.Buffer
	opts := &execext.RunCommandOptions{
		Command: fmt.Sprintf("go build -o %s/%s %s %s", TARGET_DIR, target, flagString.String(), source),
		Dir:     ".",
		Stdout:  execext.DevNull{},
		Stderr:  &stderr,
	}
	if err := execext.RunCommand(context.Background(), opts); err != nil {
		return err
	}

	if verbose > 0 || stderr.Len() > 5 {
		// go build -x output to stderr
		logger.Info(stderr.String())
	}

	if name == target {
		logger.SuccessF("Compiled %s successfully", target)
	} else {
		logger.SuccessF("Compiled %s to %s successfully", name, target)
	}

	return nil
}

func multiCompile(name, source string, buildID int) []string {
	targes := make([]string, 0, MULTI_COUNT)
	for _, goos := range OSS {
		for _, goarch := range ARCHES {
			target := fmt.Sprintf("%s-%s-%s", name, goos, goarch)
			if err := compile(name, source, target, buildID); err != nil {
				logger.Error(err)
			} else {
				targes = append(targes, target)
			}
		}
	}
	return targes
}

func archiveTargets(name string) error {
	releaseTarget := fmt.Sprintf("%s/%s", RELEASE_DIR, name)
	command := fmt.Sprintf(`
		rm -rf %s* && mkdir -p %s && \
		mv %s/%s-* %s/ && \
		cd %s && \
		shasum -a 256 %s* > checksum256 && \
		cd .. && \
		pwd && ls && \
		tar -czf %s.tar.gz %s/* && \
		rm -rf %s`,
		releaseTarget, releaseTarget,
		TARGET_DIR, name, releaseTarget,
		releaseTarget,
		name,
		name, name,
		name)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	opts := &execext.RunCommandOptions{
		Command: command,
		Dir:     ".",
		Stdout:  &stdout,
		Stderr:  &stderr,
	}
	if err := execext.RunCommand(context.Background(), opts); err != nil {
		logger.Debug(command)
		logger.Error(err)
		logger.Debug(stdout.String())
		logger.Error(stderr.String())
		return err
	}
	logger.SuccessF("Archived %s successfully", name)
	return nil
}

var rootCmd = &cobra.Command{
	Use:     "go run ./cli/_compile [flags] name...",
	Example: "  go run ./cli/_compile virzz\n  go run ./cli/_compile -R virzz jwttool",
	Short:   "Helper for building Virzz",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Clean Build Directory
		if clean {
			if err := os.RemoveAll(TARGET_DIR); err != nil {
				return err
			}
			if err := os.Mkdir(TARGET_DIR, 0775); err != nil {
				return err
			}
			logger.Success("Cleaned build directory")
			return nil
		}

		// Compile Virzz Projects
		if len(args) > 0 {
			sourceNames := make(map[string]string)
			publicNames := getPublicProjects()
			logger.Debug(publicNames)
			// all public
			if utils.SliceContains(args, "public") {
				for _, name := range publicNames {
					sourceNames[name] = fmt.Sprintf("%s/%s", PUBLIC_DIR, name)
				}
			}
			// inner public
			for _, name := range utils.Intersection(publicNames, args) {
				sourceNames[name] = fmt.Sprintf("%s/%s", PUBLIC_DIR, name)
			}
			// specific
			for _, name := range utils.Intersection(getSpecialProjects(), args) {
				sourceNames[name] = fmt.Sprintf("%s/%s", SOURCE_DIR, name)
			}

			// Compile
			for name, source := range sourceNames {
				buildID, err := BuildID.Inc(name)
				if err != nil {
					logger.Error(err)
				}
				if multi {
					logger.InfoF("Start compiling %s for all platforms", name)
					if len(multiCompile(name, source, buildID)) == 6 {
						logger.Success("All platforms compiled successfully")
						if archive {
							logger.Info("Start archiving")
							archiveTargets(name)
						}
					} else {
						logger.Error("Lost some platform binaries, try compile again")
					}

				} else {
					logger.InfoF("Start compiling %s", name)
					if err := compile(name, source, name, buildID); err != nil {
						logger.ErrorF("Compile %s fail: %v", name, err)
					}
				}
			}
			return nil
		}

		return cmd.Help()
	},
}

var (
	release bool = false
	archive bool = false
	multi   bool = false
	clean   bool = false
	force   bool = false

	verbose int = 0
)

func init() {
	rootCmd.Flags().BoolVarP(&release, "release", "R", false, "Build release version")
	rootCmd.Flags().BoolVarP(&archive, "archive", "A", false, "Archive release packages")
	rootCmd.Flags().BoolVarP(&multi, "multi", "M", false, "Compile multi-platform binaries")
	rootCmd.Flags().BoolVarP(&clean, "clean", "C", false, "Clean build files")
	rootCmd.Flags().BoolVarP(&force, "force", "F", false, "Force to Compile")

	rootCmd.Flags().CountVarP(&verbose, "verbose", "v", "Print verbose information")

}

func main() {
	logger.SetDebug(true)
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
