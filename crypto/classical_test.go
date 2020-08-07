package crypto

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
	r, _ := Morse("Virink, 123!", false)
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
