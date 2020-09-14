package collision

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// DictTable -
type DictTable struct {
	length int
	table  []byte
	prefix []byte
	suffix []byte

	result map[int]string

	collision     func(string) bool
	collisionByte func([]byte) bool

	done chan struct{}
	wg   *sync.WaitGroup
}

var (
	defaultLength = 4
	defaultTable  = []byte("abcdefghijklnmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	defaultPrefix = []byte("")
	defaultSuffix = []byte("")
)

// SetTable -
func (dt *DictTable) SetTable(table []byte) {
	dt.table = table
}

// SetLength -
func (dt *DictTable) SetLength(length int) {
	dt.length = length
}

// SetCollisionByte -
func (dt *DictTable) SetCollisionByte(fun func([]byte) bool) {
	dt.collisionByte = fun
}

// Results -
func (dt *DictTable) Results() map[int]string {
	return dt.result
}

// ProcessCollision -
func (dt *DictTable) ProcessCollision() {
	if dt.collision == nil && dt.collisionByte == nil {
		return
	}
	var helper func([]byte)
	helper = func(secret []byte) {
		select {
		case <-dt.done:
			return
		default:
			if len(secret) == dt.length {
				if dt.collisionByte(secret) {
					dt.result[0] = string(secret)
					close(dt.done)
					return
				}
			}
			if len(secret) < dt.length {
				for _, cc := range dt.table {
					helper(append(secret, cc))
				}
			}
		}
	}
	for _, ct := range dt.table {
		dt.wg.Add(1)
		go func(ct byte) {
			defer dt.wg.Done()
			helper([]byte{ct})
		}(ct)
	}
	dt.wg.Add(1)
	go func() {
		defer dt.wg.Done()
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
		for {
			select {
			case <-dt.done:
				return
			case <-interrupt:
				close(dt.done)
				return
			}
		}
	}()
	dt.wg.Wait()
}

// NewDictTable -
func NewDictTable() *DictTable {
	dt := &DictTable{}
	dt.table = defaultTable
	dt.length = defaultLength
	dt.prefix = defaultPrefix
	dt.suffix = defaultSuffix
	dt.done = make(chan struct{})
	dt.wg = &sync.WaitGroup{}
	dt.result = make(map[int]string)
	return dt
}
