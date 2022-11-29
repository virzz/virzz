package jwttool

import (
	"fmt"
	"os"
	"testing"
)

func TestPrintJWT(t *testing.T) {
	r, err := printJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
}

func TestCrackJWT(t *testing.T) {
	r, err := crackJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidmlyaW5rIn0.63La-xrRjx38xDkgrNHYfYHVgjB83bZsJMSa5luusgY", 4, 5, "abcdefghijklnmopqrstuvwxyz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "xkedd" {
		t.Fail()
	}
}

func TestModifyJWT(t *testing.T) {
	r, err := modifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidmlyaW5rIiwicm9sZSI6Imd1ZXN0In0.bPb06hMv6GA73WNOEO1D_HMyal6hS1ofBDIsRL3vszg", true, "", map[string]string{"role": "admin"}, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJuYW1lIjoidmlyaW5rIiwicm9sZSI6ImFkbWluIn0." {
		t.Fail()
	}
}
