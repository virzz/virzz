package base100

import "testing"

func TestInvalidInput(t *testing.T) {
	if _, err := Decode("aaaa"); err.Error() != "data is invalid" {
		t.Error(err)
	}
}

var tests = []struct {
	name  string
	text  string
	emoji string
}{
	{
		"ASCII",
		"hello",
		"👟👜👣👣👦",
	},
	{
		"Cyrillic",
		"РАШ Бэ",
		"📇💗📇💇📇💟🐗📇💈📈💄",
	},
	{
		"HelloUnicode",
		"Hello, 世界",
		"🐿👜👣👣👦🐣🐗📛💯💍📞💌💃",
	},
}

func TestDecode(t *testing.T) {
	for _, test := range tests {
		res, err := Decode(test.emoji)
		if err != nil {
			t.Errorf("%v: Unexpected error: %v", test.name, err)
		}
		if string(res) != test.text {
			t.Errorf("%v: Expected to get '%v', got '%v'", test.name, test.text, string(res))
		}
	}
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		res := Encode([]byte(test.text))
		if res != test.emoji {
			t.Errorf("%v: Expected to get '%v', got '%v'", test.name, test.emoji, res)
		}
	}
}

func TestFlow(t *testing.T) {
	text := []byte("the quick brown fox 😂😂👌👌👌 over the lazy dog привет")
	res, err := Decode(Encode(text))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(res) != string(text) {
		t.Errorf("Expected to get '%v', got '%v'", string(text), string(res))
	}
}
