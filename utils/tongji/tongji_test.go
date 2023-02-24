package tongji_test

import (
	"testing"

	"github.com/virzz/virzz/utils/tongji"
)

func TestTongji(t *testing.T) {
	go tongji.Tongji("http://name.tool.virzz.com", "app", "v0.1.1")
}
