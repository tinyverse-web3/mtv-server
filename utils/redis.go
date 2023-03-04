package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/gomodule/redigo/redis"
)

var redisCache cache.Cache

func InitRedis() {
	logs.Info("redis init start")
	redisHost, _ := config.String("redis::address")
	logs.Info(redisHost)
	dataBase, _ := config.String("redis::database")
	password, _ := config.String("redis::password")
	redisKey, _ := config.String("redis::key")
	config := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%s","password":"%s"}`, redisKey, redisHost, dataBase, password)
	var err error
	redisCache, err = cache.NewCache("redis", config)
	if err != nil || redisCache == nil {
		errMsg := "failed to init redis"
		logs.Error(errMsg, err)
		panic(errMsg)
	}
	logs.Info("redis init end")
}

func SetStr(key, value string, timeout time.Duration) (err error) {
	err = redisCache.Put(context.Background(), key, value, timeout)
	if err != nil {
		logs.Error("set key:", key, ",value:", value, err)
	}
	return
}

func GetStr(key string) (value string) {
	v, _ := redisCache.Get(context.Background(), key)
	value = string(v.([]byte)) //这里的转换很重要，Get返回的是interface
	return
}

func DelKey(key string) (err error) {
	err = redisCache.Delete(context.Background(), key)
	return
}
