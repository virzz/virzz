package bilibili

import (
	"fmt"
	"testing"
)

// parse

// func TestBilibili(t *testing.T) {
// 	Bilibili("https://www.bilibili.com/video/BV1Xk4y1m7N5")
// 	Bilibili("BV1Xk4y1m7N5")
// }

func TestParseVideoInfo(t *testing.T) {
	pi, err := ParseVideoInfo("BV1Xk4y1m7N5")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(pi)
}
