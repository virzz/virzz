package main

import (
	"reflect"
	"runtime"
	"time"

	"github.com/virzz/logger"
)

type StructName struct {
	FieldName uint64
	Length    uint64
}

func (t *StructName) String() string {
	return "StructName"
}
func (t *StructName) Orz() string {
	return time.Now().String()
}

func main() {
	t := new(StructName)
	t.FieldName = 1
	t.Length = 2
	logger.Info(runtime.GOROOT())
	logger.Success(t)
	logger.Success(t.Orz())
	v := reflect.ValueOf(t)
	logger.Success(v.String())
	logger.Success(v.Type().String())
	panic(t)
}
