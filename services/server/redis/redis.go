package redis

import (
	"context"
	"fmt"
	"time"

	red "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

var (
	rClient *red.Client
	once    utils.OncePlus
	ctx     = context.Background()
)

func Connect() error {
	return once.Do(func() error {

		if !viper.IsSet("redis") {
			logger.Fatal("Not set redis config")
		}

		rClient = red.NewClient(&red.Options{
			Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
			DB:       viper.GetInt("redis.db"),
			Password: "",
		})

		res := rClient.Ping(ctx)
		if res.Err() != nil {
			logger.Fatal(res.Err())
		}

		return nil
	})
}

func R() *red.Conn {
	return rClient.Conn(ctx)
}

func Debug() {
	rClient.Options().OnConnect = func(ctx context.Context, conn *red.Conn) error {
		logger.DebugF("conn=%v\n", conn)
		return nil
	}
}

// Get -
func Get(key string) (result string, err error) {
	res := R().Get(ctx, key)
	if res.Err() != nil {
		return "", err
	}
	result, err = res.Result()
	if err != nil {
		logger.Error(err)
	}
	return result, nil
}

// SetEx -
func SetEx(key, value string, expire ...int) (string, error) {
	_expire := 3600
	if len(expire) > 0 && expire[0] > 0 && expire[0] < 7*3600 {
		_expire = expire[0]
	}
	if err := R().Set(ctx, key, value, time.Duration(_expire)*time.Second).Err(); err != nil {
		return "", err
	}
	return value, nil
}

func Set(key, value string) (string, error) {
	if err := R().Set(ctx, key, value, 0).Err(); err != nil {
		return "", err
	}
	return value, nil
}
