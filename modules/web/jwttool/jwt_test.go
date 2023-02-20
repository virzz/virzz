package jwttool

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/virzz/logger"
)

func init() {
	logger.SetDebug(true)
}

func TestJWTPrint(t *testing.T) {
	r, err := JWTPrint("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
}

func TestJWTCrack(t *testing.T) {
	r, err := JWTCrack("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidmlyaW5rIn0.63La-xrRjx38xDkgrNHYfYHVgjB83bZsJMSa5luusgY", 4, 5, "abcdeijklvwxyz", "", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if r != "xkedd" {
		t.Fail()
	}
	fmt.Println(r)
}

func TestJWTModify(t *testing.T) {
	r, err := JWTModify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidmlyaW5rIiwicm9sZSI6Imd1ZXN0In0.bPb06hMv6GA73WNOEO1D_HMyal6hS1ofBDIsRL3vszg", true, "", map[string]string{"role": "admin"}, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJuYW1lIjoidmlyaW5rIiwicm9sZSI6ImFkbWluIn0." {
		t.Fail()
	}
}

func TestSpilt(t *testing.T) {
	alphabet := defaultAlphabet
	num := runtime.NumCPU()
	size := len(alphabet)/num + 1
	logger.DebugF("num = %d size = %d", num, size)
	for i := 0; i < num; i++ {
		start := i * size
		if start > len(alphabet) {
			break
		}
		end := start + size
		if end > len(alphabet) {
			end = len(alphabet)
		}
		logger.DebugF("%d S:%d E:%d %s", i, start, end, alphabet[start:end])
	}
}
