package downloader

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/virzz/logger"
)

const (
	DefaultLimit   = 10
	DefaultDelay   = 0
	DefaultTimeout = 5
)

// Downloader -
type Downloader struct {
	timeout    int64 //  超时
	limit      int64 // 并发数
	delay      int64 // 请求延迟(limit = 1)
	httpClient *http.Client
	headers    map[string]string
	workers    chan *downloadTask
	errFile    map[string]int
	results    map[string]interface{}
	result     bool
	lock       sync.Mutex
}

type downloadTask struct {
	Target   string
	DestPath string
}

// SetTimeout -
func (d *Downloader) SetTimeout(timeout int64) *Downloader {
	if timeout > 60 {
		timeout = 60
	}
	d.timeout = timeout
	return d
}

// SetLimit -
func (d *Downloader) SetLimit(limit int64) *Downloader {
	if limit > 50 {
		limit = 50
	}
	d.limit = limit
	return d
}

// SetDelay -
func (d *Downloader) SetDelay(delay int64) *Downloader {
	d.delay = delay
	if delay > 0 {
		d.limit = 1
	}
	return d
}

// SetHeader -
func (d *Downloader) SetHeader(key, value string) *Downloader {
	d.headers[key] = value
	return d
}

// SetResult -
func (d *Downloader) SetResult() *Downloader {
	d.result = !d.result
	return d
}

// AddTask -
func (d *Downloader) AddTask(target, dest string) *Downloader {
	d.workers <- &downloadTask{
		Target:   target,
		DestPath: dest,
	}
	return d
}

// AddTasks -
func (d *Downloader) AddTasks(tasks map[string]string) *Downloader {
	for target, dest := range tasks {
		d.AddTask(target, dest)
	}
	return d
}

// Init -
func (d *Downloader) Init() *Downloader {
	d.workers = make(chan *downloadTask, 1024)
	d.errFile = make(map[string]int)
	d.results = make(map[string]interface{})
	d.headers = make(map[string]string)
	d.httpClient = &http.Client{
		Timeout: time.Duration(d.timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return d
}

// Head -
func (d *Downloader) Head(work *downloadTask) (resp *http.Response, err error) {
	var (
		req *http.Request
	)
	// Head
	req, err = http.NewRequest("HEAD", work.Target, nil)
	if err != nil {
		logger.Error("Head", err)
		return
	}
	// Header
	for k, v := range d.headers {
		req.Header.Set(k, v)
	}
	resp, err = d.httpClient.Do(req)
	if err != nil {
		logger.Error("Do", err)
		return
	}
	if resp.StatusCode == 404 {
		return
	}
	return resp, nil
}

// Fetch -
func (d *Downloader) Fetch(work *downloadTask) (err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	// Fix: Traversal
	if strings.Contains(work.DestPath, "..") {
		err = fmt.Errorf("invalid path with '..': %s", work.DestPath)
		logger.Error(err)
		return
	}
	_resp, err := d.Head(work)
	if err != nil {
		logger.Error(err)
		return
	}
	size, err := strconv.ParseInt(_resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		logger.Error(err)
		return
	}
	// Already exists
	if fi, _ := os.Stat(work.DestPath); fi != nil && fi.Size() == size {
		d.results[work.DestPath] = true
		return nil
	}
	// Get
	req, err = http.NewRequest("GET", work.Target, nil)
	if err != nil {
		logger.Error("GET", err)
		return
	}
	// Header
	for k, v := range d.headers {
		req.Header.Set(k, v)
	}
	resp, err = d.httpClient.Do(req)
	if err != nil {
		logger.Error("Do", err)
		return
	}
	if resp.StatusCode == 404 {
		return
	}
	if d.result {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Body", err)
			return err
		}
		d.results[work.DestPath] = body
		return nil
	}
	err = os.MkdirAll(filepath.Dir(work.DestPath), 0700)
	if err != nil && !os.IsExist(err) {
		return
	}
	f, err := os.Create(work.DestPath)
	if err != nil {
		logger.Error("Create", err)
		return
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		logger.Error("Copy", err)
	}
	return
}

// Results -
func (d *Downloader) Results() map[string]interface{} {
	return d.results
}

// Errors -
func (d *Downloader) Errors() map[string]int {
	return d.errFile
}

// Reset -
func (d *Downloader) Reset() *Downloader {
	return d.Init()
}

// PrintResults -
func (d *Downloader) PrintResults() {
	for uri := range d.results {
		logger.Info("Fetched ", uri)
	}
	for uri := range d.errFile {
		logger.Error("Fetched ", uri)
	}
}

var (
	ErrInterrupt = fmt.Errorf("finish by interrupt")
)

// Start -
func (d *Downloader) Start() error {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	p, _ := ants.NewPoolWithFunc(int(d.limit), func(_work interface{}) {
		work := _work.(*downloadTask)
		logger.Debug("Pool Fetching ", work.Target)

		if d.delay > 0 {
			time.Sleep(time.Duration(d.delay) * time.Second)
		}

		if err := d.Fetch(work); err != nil {
			d.lock.Lock()
			num, ok := d.errFile[work.Target]
			if ok {
				if num > 2 {
					logger.Error(err.Error())
				} else {
					d.workers <- work
					d.errFile[work.Target]++
				}
			} else {
				d.errFile[work.Target] = 0
			}
			d.lock.Unlock()
		}
	})
	defer func() {
		if !p.IsClosed() {
			p.Release()
		}
	}()

	for {
		select {
		case <-interrupt:
			close(d.workers)
			p.Release()
			return ErrInterrupt
		case work, ok := <-d.workers:
			if !ok {
				break
			}
			go p.Invoke(work)
		case <-time.After(3 * time.Second):
			if p.Running() == 0 {
				return nil
			}
		}
	}
}

// NewDownloader -
func NewDownloader() *Downloader {
	d := &Downloader{}
	d.Init().SetLimit(DefaultLimit).SetDelay(DefaultDelay).SetTimeout(DefaultTimeout)
	return d
}

func SigleFetch(target, dest string) error {
	d := NewDownloader()
	err := d.Fetch(&downloadTask{Target: target, DestPath: dest})
	if err != nil {
		return err
	}
	return nil
}
