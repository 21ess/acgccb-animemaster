package db

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "password",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(Ctx, 3*time.Second)
	defer cancel()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Panicf("redis connect err: %v", err)

	}
	log.Println("redis connect success")
}
