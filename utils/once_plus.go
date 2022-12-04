package utils

import (
	"sync"
	"sync/atomic"
)

type OncePlus struct {
	m    sync.Mutex
	done uint32
}

func (o *OncePlus) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.SlowDo(f)
}

func (o *OncePlus) SlowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
