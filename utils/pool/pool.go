package pool

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/virzz/logger"
)

var (
	ErrInterrupt = fmt.Errorf("finish by interrupt")
	ErrTimeout   = fmt.Errorf("finish by timeout")
)

func Start[T comparable](pf func(T) bool, args ...T) error {
	return StartWithSize(runtime.NumCPU(), pf, args...)
}

func StartWithSize[T comparable](size int, pf func(T) bool, args ...T) error {
	closeCh := make(chan struct{}, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	p, _ := ants.NewPoolWithFunc(size, func(arg interface{}) {
		if pf(arg.(T)) {
			closeCh <- struct{}{}
			close(closeCh)
		}
	})
	defer func() {
		if !p.IsClosed() {
			p.Release()
		}
	}()

	// Invoke
	logger.Debug("Invoke...")
	if len(args) > 1 {
		for _, arg := range args {
			go p.Invoke(arg)
		}
	} else if len(args) == 1 {
		if reflect.TypeOf(args[0]).Kind() == reflect.Chan {
			v := reflect.ValueOf(args[0])
			go func() {
				for {
					if v, ok := v.Recv(); ok {
						p.Invoke(v)
					}
				}
			}()
		} else {
			p.Invoke(args[0])
		}
	} else {
		for i := 0; i < size; i++ {
			go p.Invoke(nil)
		}
	}

	// Listening
	go func() {
		for {
			time.Sleep(1 * time.Second)
			logger.DebugF("Listening: R:%d W:%d", p.Running(), p.Waiting())
			if p.IsClosed() {
				return
			}
			if p.Running() == 0 && p.Waiting() == 0 {
				closeCh <- struct{}{}
				close(closeCh)
				logger.DebugF("Listening: R:%d W:%d", p.Running(), p.Waiting())
				return
			}
		}
	}()

	// Waiting
	logger.Debug("Waiting...")
	select {
	case <-interrupt:
		p.Release()
		return ErrInterrupt
	case <-closeCh:
		p.Release()
		return nil
	}
}
