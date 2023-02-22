package parser

import (
	"fmt"
	"testing"
)

func TestParseProcNetTcp(t *testing.T) {
	r, err := ParseProcNet("../../tests/proc_net_tcp")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
func TestParseProcNetUdp(t *testing.T) {
	r, err := ParseProcNet("../../tests/proc_net_udp")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
