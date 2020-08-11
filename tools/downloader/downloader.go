package downloader

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/virink/virzz/common"
)

// Downloader -
type Downloader struct {
	httpClient *http.Client
	timeout    int64 //  超时
	limit      int64 // 并发数
	delay      int64 // 请求延迟(limit = 1)
	workers    chan downloadTask
	cancelChan chan bool
	errFile    map[string]int
	results    map[string]interface{}
	result     bool
}

type downloadTask struct {
	Target   string
	DestPath string
}

// SetTimetou -
func (d *Downloader) SetTimetou(timeout int64) *Downloader {
	if timeout > 30 {
		timeout = 30
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

// SetResult -
func (d *Downloader) SetResult() *Downloader {
	d.result = !d.result
	return d
}

// AddTask -
func (d *Downloader) AddTask(target, dest string) *Downloader {
	d.workers <- downloadTask{
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
	d.workers = make(chan downloadTask, 102400)
	d.errFile = make(map[string]int)
	d.results = make(map[string]interface{})
	return d
}

// Fetch -
func (d *Downloader) Fetch(work downloadTask) (err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest("GET", work.Target, nil)
	if err != nil {
		common.Logger.Error("NewRequest", err)
		return
	}
	resp, err = d.httpClient.Do(req)
	if err != nil {
		common.Logger.Error("Do", err)
		return
	}
	if resp.StatusCode == 404 {
		return
	}
	if d.result {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			common.Logger.Error("Body", err)
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
		common.Logger.Error("Create", err)
		return
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		common.Logger.Error("Copy", err)
	}
	return
}

// Results -
func (d *Downloader) Results() map[string]interface{} {
	return d.results
}

// Reset -
func (d *Downloader) Reset() *Downloader {
	return d.Init()
}

// Start -
func (d *Downloader) Start() error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	d.httpClient = &http.Client{
		Timeout: time.Duration(d.timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}

	wg := &sync.WaitGroup{}
	d.cancelChan = make(chan bool)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	for i := int64(0); i < d.limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-d.cancelChan:
					return
				case <-time.After(3 * time.Second):
					defer func() {
						if recover() != nil {
						}
					}()
					close(d.cancelChan)
				case work := <-d.workers:
					common.Logger.Debug("Downloder Fetch", work.Target)
					// Fetch
					if err := d.Fetch(work); err != nil {
						common.Logger.Debug(err.Error())
						num, ok := d.errFile[work.Target]
						if ok {
							if num > 2 {
								common.Logger.Error(err.Error())
							} else {
								d.workers <- work
								d.errFile[work.Target]++
							}
						} else {
							d.errFile[work.Target] = 0
						}
					}
					// delay
					if d.delay > 0 {
						time.Sleep(time.Duration(d.delay) * time.Second)
					}
				}
			}
		}()
	}

	wg.Wait()

	return nil
}

// NewDownloader -
func NewDownloader() *Downloader {
	d := &Downloader{}
	d.Init().SetLimit(10).SetDelay(0).SetTimetou(3)
	return d
}
