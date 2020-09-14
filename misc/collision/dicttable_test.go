package collision

import (
	"bytes"
	"testing"
)

func TestDT(t *testing.T) {
	dt := NewDictTable()
	dt.SetTable([]byte("abcdefghijklnmopqrstuvwxzy"))
	dt.SetLength(4)
	dt.SetCollisionByte(func(secret []byte) bool {
		// fmt.Println(string(secret))
		return bytes.Equal(secret, []byte("bdcd"))
	})
	dt.ProcessCollision()
}
