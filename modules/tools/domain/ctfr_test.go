package domain

import (
	"fmt"
	"testing"
)

func TestCtfr(t *testing.T) {
	s, err := ctfr("virzz.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
}
