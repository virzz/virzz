package parser

import (
	"fmt"
	"testing"
)

func TestParseProcNetTcp(t *testing.T) {
	r, err := parseProcNetTcp("../../tests/proc_net_tcp")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
