package hash

import (
	"fmt"
	"testing"
)

var (
	_passwd = "aewgvrweasgw"
	_hashed = "$2a$10$dmBwjxKqfD2T4n/pAaaQ2ePNHFp8U9GMes5XNfKUC8ssezx2y/2Ci"
)

func TestBcryptGenerate(t *testing.T) {
	r, err := BcryptGenerate(_passwd, 10)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}

func TestBcryptCompare(t *testing.T) {
	err := BcryptCompare(_hashed, _passwd)
	if err != nil {
		t.Fatal(err)
	}

}
