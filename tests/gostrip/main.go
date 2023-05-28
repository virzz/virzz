package main

import (
	"reflect"
	"runtime"
	"time"

	"github.com/virzz/logger"
)

type Tttt struct {
	Offset uint64
	Length uint64
}

func (t *Tttt) String() string {
	return "Tttt"
}
func (t *Tttt) Orz() string {
	return time.Now().String()
}

func main() {
	t := new(Tttt)
	t.Offset = 1
	t.Length = 2
	logger.Info(runtime.GOROOT())
	logger.Success(t)
	logger.Success(t.Orz())
	v := reflect.ValueOf(t)
	logger.Success(v.String())
	logger.Success(v.Type().String())
	panic(t)
}
