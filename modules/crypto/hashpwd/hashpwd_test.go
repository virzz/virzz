package hashpwd

import (
	"testing"

	"github.com/virzz/logger"
)

const dic = "../../../tests/hashpwd/top1k.dic"

func TestGen(t *testing.T) {
	pwd := "../../../tests/hashpwd/top1k.txt"
	if err := GenerateHashDict(pwd, dic); err != nil {
		t.Error(err)
	}
}

var testData = map[string]string{
	"qwert":    "a384b6463fc216a5f8ecb6670f86456a",                                 // md5
	"money":    "b18ffcc59b8e8aa36a4ce3968b6d8e2b",                                 // sha1-md5
	"123456":   "7c4a8d09ca3762af61e59520943dc26494f8941b",                         // sha1
	"iloveyou": "cfbf459d9d6057bc2a85477a38327b96f06b1597",                         // sha1-sha1
	"fuckyou":  "6161b0a284159565a0f7d5df2dd2698b5f87906cd91ff5322caf179b451f5a41", // sha256
	"1q2w3e":   "*5194866966b75566",                                                // mysql
	"a123456":  "f40460fe1ceec6f6785997f3319553bb",                                 // NTLM
}

func TestLookup(t *testing.T) {
	for k, v := range testData {
		r, err := LookupHashDict(dic, v)
		if err != nil {
			t.Fatal(err)
		}
		if r != k {
			logger.ErrorF("%s != %s - ", r, k, v)
		}
		logger.SuccessF("Found: %s By %s", r, v)
		// time.Sleep(1 * time.Second)
	}
}
