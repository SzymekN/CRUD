package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func GetRDB() *redis.Client {
	return RDB
}

func SetupRedisConnection() *redis.Client {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := RDB.Ping(context.Background()).Result()

	fmt.Println(pong, err)

	return RDB
}
