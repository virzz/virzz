package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/virzz/logger"

	_ "github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/execext"
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
		if version == "" {
			if name == "virzz" || name == "platform" {
				version = getVersion("p")
			} else if name != "public" {
				version = getVersion("v")
			} else {
				version = "unknown"
			}
		}
		flags["-trimpath"] = ""
		flags["-tags"] = "release"
		flags["-ldflags"] = fmt.Sprintf("-s -w -X main.BuildID=%d -X main.Version=%s", buildID, version)
	} else {
		flags["-tags"] = "debug"
		flags["-ldflags"] = fmt.Sprintf("-X main.BuildID=%d", buildID)
	}

	if len(goTags) > 0 {
		flags["-tags"] = strings.Join(append(goTags, flags["-tags"]), ",")
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

	var env = os.Environ()

	logger.Debug(name, target)

	// Multi-platform
	if name != target {
		ts := strings.Split(target, "-")
		env = append(env, "GOOS="+ts[1], "GOARCH="+ts[2], "CGO_ENABLED=0")
	}

	outputTarget := path.Join(output, target)
	buildCmd := fmt.Sprintf("%s build -o %s %s %s", builder, outputTarget, flagString.String(), source)
	var stderr bytes.Buffer
	logger.Warn("Build CMD: ", buildCmd)
	opts := &execext.RunCommandOptions{
		Command: buildCmd,
		Dir:     ".",
		Env:     env,
		Stdout:  execext.DevNull{},
		Stderr:  &stderr,
	}
	err := execext.RunCommand(context.Background(), opts)
	if verbose > 0 && stderr.Len() > 5 {
		// go build -x output to stderr
		logger.Error(stderr.String())
	}
	if err != nil {
		return err
	}

	logger.SuccessF("Compiled %s to %s successfully", name, outputTarget)

	return nil
}

func multiCompile(name, source string, buildID int) []string {
	targets := make([]string, 0, MULTI_COUNT)
	for _, goos := range OSS {
		for _, goarch := range ARCHES {
			target := fmt.Sprintf("%s-%s-%s", name, goos, goarch)
			if err := compile(name, source, target, buildID); err != nil {
				logger.Error(err)
			} else {
				targets = append(targets, target)
			}
		}
	}
	return targets
}

func archiveTargets(name string) error {
	releaseTarget := fmt.Sprintf("%s/%s", RELEASE_DIR, name)
	command := `
		rm -rf ${RELEASE}* && mkdir -p ${RELEASE} && \
		mv ${TARGET}/${NAME}-* ${RELEASE}/ && \
		cd ${RELEASE} && \
		shasum -a 256 ${NAME}* > checksum256 && \
		if [ -n "$TOGETHER" ]; then \
			tar -czf ../${NAME}.tar.gz ./* ; \
			cd .. && rm -rf ${NAME} ; \
		else \
			for f in $(ls ${NAME}*); do \
				tar -czf ${f}.tar.gz $f; \
				rm ${f}; \
			done; \
			shasum -a 256 ${NAME}* >> checksum256
		fi
		`
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var env = append(os.Environ(),
		"RELEASE="+releaseTarget,
		"NAME="+name,
		"TARGET="+output,
	)
	if archiveTogether {
		env = append(env, "TOGETHER=1")
	}
	opts := &execext.RunCommandOptions{
		Env:     env,
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
							if archiveTogether {
								logger.Info("Start archiving togher")
							} else {
								logger.Info("Start archiving")
							}
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
	release         bool
	archive         bool
	archiveTogether bool
	multi           bool
	clean           bool
	force           bool
	output          string
	goTags          []string
	verbose         int = 0
	version         string
	builder         string
)

func init() {
	rootCmd.Flags().BoolVarP(&release, "release", "R", false, "Build release version")
	rootCmd.Flags().BoolVarP(&archive, "archive", "A", false, "Archive release packages")
	rootCmd.Flags().BoolVarP(&archiveTogether, "together", "t", false, "Archive binary files in one package")
	rootCmd.Flags().BoolVarP(&multi, "multi", "M", false, "Compile multi-platform binaries")
	rootCmd.Flags().BoolVarP(&clean, "clean", "C", false, "Clean build files")
	rootCmd.Flags().BoolVarP(&force, "force", "F", false, "Force to Compile")
	rootCmd.Flags().StringVarP(&version, "version", "V", "", "Custom Build version")
	rootCmd.Flags().StringVarP(&output, "output", "o", TARGET_DIR, "Custom output path")
	rootCmd.Flags().StringSliceVarP(&goTags, "tags", "T", []string{}, "Append go build tags")
	rootCmd.Flags().StringVarP(&builder, "builder", "B", "go", "Replace `go`")
	rootCmd.Flags().CountVarP(&verbose, "verbose", "v", "Print verbose information")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
