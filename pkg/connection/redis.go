package connection

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func RedisConnection()  {

	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
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
	RedisConnection()
	return Client
}
