package basex

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/virzz/virzz/utils"
)

type (
	pair struct {
		decoded, encoded string
	}
	funx struct {
		encode, decode func(string) (string, error)
	}
)

var (
	samples = []pair{}
	funcs   = map[string]funx{
		"base16":  funx{Base16Encode, Base16Decode},
		"base36":  funx{Base36Encode, Base36Decode},
		"base62":  funx{Base62Encode, Base62Decode},
		"base85":  funx{Base85Encode, Base85Decode},
		"base91":  funx{Base91Encode, Base91Decode},
		"base92":  funx{Base92Encode, Base92Decode},
		"base100": funx{Base100Encode, Base100Decode},
	}
)

func TestBaseX(t *testing.T) {
	for i := 0; i < 20; i++ {
		samples = append(samples, pair{
			utils.RandomStringByLength(rand.Intn(50) + 20), "",
		})
	}
	for fn, fp := range funcs {
		for i, sample := range samples {
			encoded, err := fp.encode(sample.decoded)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			sample.encoded = encoded
			decoded, err := fp.decode(encoded)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			if decoded != sample.decoded {
				t.Fail()
			}
			fmt.Printf("%s - %d\nde: %s\nen: %s\n", fn, i, decoded, encoded)
		}
	}
}

func BenchmarkBaseX(b *testing.B) {
	var data = utils.RandomStringByLength(rand.Intn(50) + 20)
	var encode string
	for fn, fp := range funcs {
		b.Run(fmt.Sprintf("%sencode", fn), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fp.encode(data)
			}
		})
		encode, _ = fp.encode(data)
		b.Run(fmt.Sprintf("%sdecode", fn), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fp.decode(encode)
			}
		})
	}

}

func TestBase64Encode(t *testing.T) {
	r, err := Base64Encode("abcdefg!@#$%^&*()_+<>?{}|:", true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "YWJjZGVmZyFAIyQlXiYqKClfKzw-P3t9fDo=" {
		t.Fail()
	}
}

func TestBase64Decode(t *testing.T) {
	// YWJjZGVmZyFAIyQlXiYqKClfKzw+P3t9fDo=
	// YWJjZGVmZyFAIyQlXiYqKClfKzw-P3t9fDo=
	r, err := Base64Decode("YWJjZGVmZyFAIyQlXiYqKClfKzw-P3t9fDo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "abcdefg!@#$%^&*()_+<>?{}|:" {
		t.Fail()
	}
}

func TestBase32Encode(t *testing.T) {
	r, err := Base32Encode("abcdefg!@#$%^&*()_+<>?{}|:")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "MFRGGZDFMZTSCQBDEQSV4JRKFAUV6KZ4HY7XW7L4HI======" {
		t.Fail()
	}
}

func TestBase32Decode(t *testing.T) {
	r, err := Base32Decode("MFRGGZDFMZTSCQBDEQSV4JRKFAUV6KZ4HY7XW7L4HI")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "abcdefg!@#$%^&*()_+<>?{}|:" {
		t.Fail()
	}
}

func TestBase58Encode(t *testing.T) {
	r, err := Base58Encode("test_base58_string", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	r, err = Base58Encode("test_base58_string", "flickr")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	r, err = Base58Encode("test_base58_string", "ripple")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
}

func TestBase58Decode(t *testing.T) {
	r, err := Base58Decode("5q1dAkvfMPRxpkkujHtkssust", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println([]byte(r))
	fmt.Println(r)
	if r != "test_base58_string" {
		t.Fail()
	}
}
