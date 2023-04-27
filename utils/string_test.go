package utils

import "testing"

func TestGenerateAlphabet(t *testing.T) {
	t.Log(string(GenerateAlphabet(`\d`)))
	t.Log(string(GenerateAlphabet(`a-z`)))
	t.Log(string(GenerateAlphabet(`a-z0-9`)))
	t.Log(string(GenerateAlphabet(`a-f0-9`)))
	t.Log(string(GenerateAlphabet(`\w`)))
	t.Log(string(GenerateAlphabet(`\S`)))
}
