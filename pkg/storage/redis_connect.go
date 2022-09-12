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
		Password: "",
		DB:       0,
	})
	fmt.Println(os.Getenv("REDIS_HOST"))
	fmt.Println(os.Getenv("REDIS_HOST") + ":6379")
	fmt.Println(RDB.Options().Addr)
	fmt.Println(RDB)

	pong, err := RDB.Ping(context.Background()).Result()

	fmt.Println(pong, err)

	return RDB
}
