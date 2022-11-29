package classical

import (
	"fmt"
	"testing"
)

func TestInnerCaesar(t *testing.T) {
	var i rune
	for i = 0; i < 26; i++ {
		fmt.Println(_caesar("Hello, I'm Virink!", i))
	}
}

func TestCaesar(t *testing.T) {
	r, _ := Caesar("Hello, I'm Virink!")
	fmt.Println(r)
}

func TestRot13(t *testing.T) {
	r, _ := Rot13("Hello, I'm Virink!")
	fmt.Println(r)
	if r != "Hryyb, I'z Vvevax!" {
		t.Fail()
	}
	r, _ = Rot13(r)
	fmt.Println(r)
	if r != "Hello, I'm Virink!" {
		t.Fail()
	}
}

func TestMorse(t *testing.T) {
	r, _ := Morse("VirinK, 123!", false)
	fmt.Println(r)
	if r != "...-/../.-./../-./-.-/--..--/......../.----/..---/...--/-.-.--" {
		t.Fail()
	}
	r, _ = Morse(r, true)
	fmt.Println(r)
	if r != "Virink,[ERROR]123!" {
		t.Fail()
	}
}

func TestAtbash(t *testing.T) {
	r, _ := Atbash("svool, r'n erirmp!")
	fmt.Println(r)
	if r != "hello, i'm virink!" {
		t.Fail()
	}
}

func TestPeigen(t *testing.T) {
	r, _ := Peigen("aabbbaabaaababbababbabbbababbaabbbabaaabababbaaabb")
	fmt.Println(r)
	if r != "helloworld" {
		t.Fail()
	}
	r, _ = Peigen("thepeigenisgood")
	fmt.Println(r)
	if r != "baabbaabbbaabaaabbbbaabaaabaaaaabbaaabaaabbababaaabaabaaabbaabbbaabbbaaaabb" {
		t.Fail()
	}
}

func TestVigenere(t *testing.T) {
	r, _ := Vigenere("Mozhu is good", "secret")
	fmt.Println(r)
	// ESBYYBKKQFH
	if r != "ESBYYBKKQFH" {
		t.Fail()
	}
	r, _ = Vigenere("ESBYYBKKQFH", "secret", true)
	fmt.Println(r)
	if r != "MOZHUISGOOD" {
		t.Fail()
	}
}
