package hashpwd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/pkg/errors"
)

type SyncWriter struct {
	m      sync.Mutex
	Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(b)
}

type Glimit struct {
	n int
	c chan struct{}
}

// initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func(string), p string) {
	g.c <- struct{}{}
	go func() {
		f(p)
		<-g.c
	}()
}

func GenerateHashDict(fpath, output string) error {
	in, err := os.Open(fpath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer in.Close()

	out, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer out.Close()

	writeOut := &SyncWriter{sync.Mutex{}, out}
	wg := sync.WaitGroup{}
	g := New(8)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		wg.Add(1)
		goFunc := func(p string) {
			fmt.Fprintln(writeOut, Encrypt(p))
			wg.Done()
		}
		g.Run(goFunc, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return errors.WithStack(err)
	}
	wg.Wait()
	return nil
}

// FEAT: 逐行读取文件，一次查询多个密码？
// 目前每查询一次密码就检索整个文件
func LookupHashDict(fpath, pwd string) (string, error) {
	in, err := os.Open(fpath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if LookUp(scanner.Text(), pwd) {
			return GetPlaintext(scanner.Text()), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", errors.WithStack(err)
	}
	return "", nil
}
