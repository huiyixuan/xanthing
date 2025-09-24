package service

import (
	"fmt"
	"xanthing/config"

	"github.com/go-redis/redis"
	"github.com/spf13/cast"
)

type RedisS struct{}

var Rdb *redis.Client

func (RedisS) Init() {
	conf := config.GetConfig("redis").(map[string]any)

	Rdb = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%v", conf["addr"]),
		Password:   cast.ToString(conf["auth"]),
		DB:         cast.ToInt(conf["db"]),
		MaxRetries: 3,
	})
	_, err := Rdb.Ping().Result()
	if err != nil {
		panic("pong fail")
	}

}
