package redisclient

import (
	"fmt"
	"github.com/go-redis/redis"
)

// RedisClient redis实例
var RedisClient *redis.Client 

// InitRedis redis初始化
func InitRedis() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,     
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	} else {
		fmt.Println(pong, err)
		RedisClient = client
	}
	return client
}
