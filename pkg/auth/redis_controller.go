package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

var RDB *redis.Client
var ctx = context.Background()

func GetRDB() *redis.Client {
	return RDB
}

func SetupRedisConnection() *redis.Client {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(os.Getenv("REDIS_HOST"))
	fmt.Println(os.Getenv("REDIS_HOST") + ":6379")
	fmt.Println(RDB.Options().Addr)
	fmt.Println(RDB)

	pong, err := RDB.Ping(ctx).Result()

	fmt.Println(pong, err)

	err = RDB.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	return RDB
}

func SetToken(value string, expireTime time.Duration) error {
	rdb := GetRDB()
	fmt.Println("TUTAJ", rdb.Options().Addr)
	fmt.Println("TUTAJ", RDB.Options().Addr)

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	err = rdb.Set(ctx, "key2", "value2", 0).Err()
	if err != nil {
		panic(err)
	}

	res, err := rdb.Get(ctx, "asd").Result()

	fmt.Println(err.Error())
	fmt.Println(res)
	return nil
}
