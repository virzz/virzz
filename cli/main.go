package main

import (
	"fmt"
	"runtime"

	"github.com/virink/virzz/cli/cmd"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("本程序不支持辣鸡 windows")
		return
	}
	cmd.Execute()
}
