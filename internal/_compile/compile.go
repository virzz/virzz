package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils/execext"
)

const (
	VerboseV = 1
	VerboseX = 2
)

var (
	force   = false
	verbose = 0
	release = false

	gitRevision = ""

	goBuilder = "go"
	goOutput  = ""
	goVersion = ""
	goTags    = []string{}
)

func compile(name, source, target string) error {

	flags := make(map[string]string)

	if force {
		flags["-a"] = ""
	}

	if verbose == VerboseV {
		flags["-v"] = ""
	} else if verbose == VerboseX {
		flags["-x"] = ""
	}

	if release {
		if goVersion == "" {
			if name == "virzz" || name == "platform" {
				goVersion = getVersion("p")
			} else if name != "public" {
				goVersion = getVersion("v")
			} else {
				goVersion = "unknown"
			}
		}
		flags["-trimpath"] = ""
		flags["-tags"] = "release"
		flags["-ldflags"] = fmt.Sprintf(
			"-s -w -X main.Version=%s -X main.Revision=%s", goVersion, gitRevision)
	} else {
		flags["-tags"] = "debug"
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

	// Multi-platform
	if name != target {
		ts := strings.Split(target, "-")
		tsLen := len(ts)
		env = append(env, "GOOS="+ts[tsLen-2], "GOARCH="+ts[tsLen-1])
	}

	outputTarget := path.Join(goOutput, target)
	buildCmd := fmt.Sprintf("%s build -o %s %s %s", goBuilder, outputTarget, flagString.String(), source)
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

	return nil
}

func multiCompile(name, source string) []string {
	targes := make([]string, 0, MULTI_COUNT)
	for _, goos := range OSS {
		for _, goarch := range ARCHES {
			target := fmt.Sprintf("%s-%s-%s", name, goos, goarch)
			if err := compile(name, source, target); err != nil {
				logger.Error(err)
			} else {
				targes = append(targes, target)
			}
		}
	}
	return targes
}
