package jwttool

import (
	"fmt"
	"testing"
)

func TestCrackJWTV2(t *testing.T) {
	alphabet := []byte("abcdefghijklnmopqrstuvwxyz")
	jwts := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidmlyaW5rIn0.63La-xrRjx38xDkgrNHYfYHVgjB83bZsJMSa5luusgY"
	res, err := crackJWTV2(jwts, 4, 5, alphabet, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
