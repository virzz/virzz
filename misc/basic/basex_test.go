package basic

import (
	"fmt"
	"os"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	r, err := Base64Encode("abcdefg!@#$%^&*()_+<>?{}|:", true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "YWJjZGVmZyFAIyQlXiYqKClfKzw+P3t9fDo=" {
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
	r, err := Base58Encode("test_base58_string")
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
	r, err := Base58Decode("5q1dAkvfMPRxpkkujHtkssust")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println([]byte(r))
	fmt.Println(r)
	if r != "test_base58_string" {
		t.Fail()
	}
}
