package base92

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	bin  []byte
}{
	{"nil", nil},
	{"empty", []byte{}},
	{"zero", []byte{0}},
	{"one", []byte{1}},
	{"two", []byte{2}},
	{"ten", []byte{10}},
	{"2zeros", []byte{0, 0}},
	{"2ones", []byte{1, 1}},
	{"64zeros", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	{"65zeros", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	{"ascii", []byte("c'est une longue chanson")},
	{"utf8", []byte("Garçon, un café très fort !")},
}

func TestEncode(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			str := StdEncoding.EncodeToString(c.bin)

			ni := len(c.bin)
			if ni > 70 {
				ni = 70 // print max the first 70 bytes
			}
			na := len(str)
			if na > 70 {
				na = 70 // print max the first 70 characters
			}
			t.Logf("bin len=%d [:%d]=%v", len(c.bin), ni, c.bin[:ni])
			t.Logf("str len=%d [:%d]=%q", len(str), na, str[:na])

			got, err := StdEncoding.DecodeString(str)
			if err != nil {
				t.Errorf("Decode() error = %v", err)
				return
			}

			ng := len(got)
			if ng > 70 {
				ng = 70 // print max the first 70 bytes
			}
			t.Logf("got len=%d [:%d]=%v", len(got), ng, got[:ng])

			if (len(got) == 0) && (len(c.bin) == 0) {
				return
			}

			if !reflect.DeepEqual(got, c.bin) {
				t.Errorf("Decode() = %v, want %v", got, c.bin)
			}
		})
	}
}
