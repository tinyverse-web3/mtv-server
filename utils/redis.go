package utils

import (
	"context"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx context.Context

func InitRedis() {
	logs.Info("redis init start")
	address, _ := config.String("redis::address")
	dataBase, _ := config.Int("redis::database")
	password, _ := config.String("redis::password")

	ctx = context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       dataBase,
	})

	logs.Info("redis init end")
}

func SetStr(key, value string, timeout time.Duration) (err error) {
	err = rdb.Set(ctx, key, value, timeout).Err()
	if err != nil {
		logs.Error(err)
	}
	return
}

func GetStr(key string) (value string, err error) {
	value, err = rdb.Get(ctx, key).Result()
	if err != nil {
		logs.Error(err)
	}
	return
}

func DelKey(key string) (err error) {
	// err = redisCache.Delete(context.Background(), key)
	return
}
