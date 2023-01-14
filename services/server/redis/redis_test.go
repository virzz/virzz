package redis

import (
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestConn(t *testing.T) {
	testValue := "argwragerahbetd"
	viper.Set("redis", map[string]interface{}{
		"host": "127.0.0.1",
		"port": 6379,
		"db":   0,
	})
	if err := Connect(); err != nil {
		t.Fatal(err)
	}
	if _, err := SetEx("test", testValue, 2); err != nil {
		t.Fatal(err)
	}
	res, err := Get("test")
	if err != nil {
		t.Fatal(err)
	}
	if res != testValue {
		t.Fatal(res)
	}
	time.Sleep(time.Duration(3) * time.Second)
	res, err = Get("test")
	if err != nil {
		t.Fatal(err)
	}
	if res != "" {
		t.Fatal(res)
	}
}
