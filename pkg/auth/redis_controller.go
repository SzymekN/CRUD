package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/go-redis/redis/v8"
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

	return RDB
}

func getSigningKey(token string) (string, error) {
	rdb := GetRDB()

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	c := context.Background()
	res, err := rdb.Get(c, token).Result()

	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading token:"+err.Error())
		fmt.Println("ERROR reading token:", err.Error())
		return "", err

	}

	return res, nil
}

func SetToken(token, key string, expireTime time.Duration) error {
	rdb := GetRDB()

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	err = rdb.Set(ctx, token, key, 0).Err()
	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing token:"+err.Error())
		panic(err)
	}

	return nil
}
