//go:build go1.20
// +build go1.20

package utils

import (
	"unsafe"
)

// S2B converts string to a byte slice without memory allocation.
func S2B(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
