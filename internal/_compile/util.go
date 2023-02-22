package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils/execext"
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
