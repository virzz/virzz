package parser

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gehaxelt/ds_store"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/downloader"
)

const dsStoreFile = ".DS_Store"

func parseDSStore(data []byte) ([]string, error) {
	allocator, err := ds_store.NewAllocator(data)
	if err != nil {
		return nil, err
	}
	fs, err := allocator.TraverseFromRootNode()
	if err != nil {
		return nil, err
	}
	var m = make(map[string]struct{}, 0)
	for _, f := range fs {
		m[f] = struct{}{}
	}
	fs = make([]string, 0, len(m))
	for k := range m {
		fs = append(fs, fmt.Sprintf("- %s", k))
	}
	sort.Strings(fs)
	return fs, nil
}

func urlJoin(base string, v ...interface{}) string {
	for i, x := range v {
		v[i] = strings.Trim(x.(string), "/")
	}
	return strings.Trim(base, "/") + strings.ReplaceAll(fmt.Sprintf(strings.Repeat("/%s", len(v)), v...), "//", "/")
}

func fetchAndParseDSStore(base string, uris []string, client *downloader.Downloader, fs *[]string) {
	for _, uri := range uris {
		// client.AddTask(fmt.Sprintf("%s%s/%s", base, uri, dsStoreFile), uri)
		client.AddTask(urlJoin(base, uri, dsStoreFile), uri)
	}
	client.Start()
	res := client.Results()
	nextUris := make([]string, 0)
	for uri, body := range res {
		// logger.Success("Fetched ", fmt.Sprintf("%s/%s", uri, dsStoreFile))
		logger.Success("Fetched ", urlJoin(uri, dsStoreFile))
		_fs, err := parseDSStore(body.([]byte))
		if err != nil {
			continue
		}
		// add to results
		*fs = append(*fs, _fs...)
		// new task
		nu := make([]string, len(_fs))
		for i, f := range _fs {
			// nu[i] = fmt.Sprintf("%s/%s", uri, f)
			nu[i] = urlJoin(uri, f)
		}
		nextUris = append(nextUris, nu...)
	}
	if len(nextUris) > 0 {
		// TODO: Download found files
		logger.Debug("nextUris", nextUris)
		fetchAndParseDSStore(base, nextUris, client.Reset(), fs)
	}
}

func DSStore(s string, download bool) (string, error) {
	var data []byte
	// Web
	if err := utils.ValidArg(s, "url"); err == nil && strings.HasPrefix(s, "http") {
		client := downloader.NewDownloader().SetLimit(1).SetResult()
		fs := make([]string, 0)
		s = strings.TrimSuffix(s, dsStoreFile)
		_u, err := url.Parse(s)
		if err != nil {
			return "", err
		}
		base := fmt.Sprintf("%s://%s", _u.Scheme, _u.Host)
		logger.Normal("Fetch URL", base)
		fetchAndParseDSStore(base, []string{_u.Path}, client, &fs)
		res := make([]string, len(fs))
		for i, f := range fs {
			res[i] = fmt.Sprintf("%s/%s", base, f)
		}
		return strings.Join(res, "\n"), err
	}
	// Local File
	if !strings.HasSuffix(s, dsStoreFile) {
		s = filepath.Join(s, dsStoreFile)
	}
	logger.Success("Fetch File: ", s)
	_, err := os.Stat(s)
	if !os.IsNotExist(err) {
		data, err = os.ReadFile(s)
		if err != nil {
			return "", err
		}
		fs, err := parseDSStore(data)
		return strings.Join(fs, "\n"), err
	}
	return "", err
}
