package connection

import (
	"context"
	"fmt"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func RedisConnection() {

	Client = redis.NewClient(&redis.Options{
		Addr:     config.LocalConfig.REDIS_HOST + ":" + config.LocalConfig.REDIS_PORT,
		Password: config.LocalConfig.REDIS_PASS,
		DB:       0,
	})
	red, err := Client.Ping(context.Background()).Result()
	fmt.Println(red)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("redis connection successful...")
}

func Redis() *redis.Client {
	if Client == nil {
		RedisConnection()
	}
	return Client
}
