package nettool

import (
	"fmt"
	"testing"
)

func TestMACToDec(t *testing.T) {
	r, err := Mac2Dec("00:00:5e:00:53:01")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "1577079553" {
		t.Fail()
	}
}

func TestDec2MAC(t *testing.T) {
	r, err := Dec2MAC("1577079553")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "00:00:5e:00:53:01" {
		t.Fail()
	}
}
