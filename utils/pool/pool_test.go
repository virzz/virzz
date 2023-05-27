package pool_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils/pool"
)

func init() {
	logger.SetDevFlags()
}

var (
	done    = make(chan struct{}, 1)
	argFunc = func(arg string) bool {
		time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)
		logger.Debug("do argFunc = ", arg)
		if rand.Intn(10) == 9 {
			logger.Debug("done start")
			done <- struct{}{}
			logger.Debug("done end")
			close(done)
			return true
		}
		return false
	}
)

func TestSingle(t *testing.T) {
	pool.Start(argFunc, "test")
}
