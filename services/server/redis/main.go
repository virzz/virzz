package models

import (
	"fmt"
	"time"

	red "github.com/go-redis/redis/v8"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils"
)

// RedisPool Redis Pool
type RedisPool struct {
	pool *red.Pool
}

var (
	Redis *RedisPool
	once  utils.OncePlus
	conf  = common.Conf.Redis
)

// var scriptGetDel = red.NewScript(1, `
// 	local r = redis.call('get', KEYS[1])
// 	if (r) then
// 		redis.call('del', KEYS[1])
// 	end
//     return r
// `)

func Connect() error {
	return once.Do(func() error {
		Redis = new(RedisPool)
		Redis.pool = &red.Pool{
			MaxIdle:     256,
			MaxActive:   0,
			IdleTimeout: time.Duration(120),
			Dial: func() (red.Conn, error) {
				return red.Dial(
					"tcp",
					fmt.Sprintf("%s:%d", conf.Host, conf.Port),
					red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
					red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
					red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
					red.DialDatabase(conf.Db),
					red.DialPassword(conf.Pass),
				)
			},
		}
		return nil
	})
}

// Get -
func Get(key string) (result string, err error) {
	conn := Redis.pool.Get()
	if err = conn.Err(); err != nil {
		return "", err
	}
	defer conn.Close()
	if result, err = red.String(conn.Do("get", key)); err != nil {
		return "", err
	}
	return result, nil
}

// Set -
func Set(key, value string, expire ...int) (result string, err error) {
	conn := Redis.pool.Get()
	if err = conn.Err(); err != nil {
		return "", err
	}
	defer conn.Close()
	_expire := 3600
	if len(expire) > 0 && expire[0] > 60 && expire[0] < 7*3600 {
		_expire = expire[0]
	}
	if _, err = conn.Do("setex", key, _expire, value); err != nil {
		return "", err
	}
	return value, nil
}
