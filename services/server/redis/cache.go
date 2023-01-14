package redis

import (
	"time"

	"github.com/go-redis/cache/v8"
)

var Cache *cache.Cache

func InitCache() {
	Cache = cache.New(&cache.Options{
		Redis:      rClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
}
