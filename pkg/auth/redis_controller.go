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

func getSigningKey() (string, error) {
	rdb := GetRDB()

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	res, err := rdb.Get(ctx, "key").Result()

	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading key:"+err.Error())
		fmt.Println("ERROR reading key:", err.Error())
		return "", err

	}

	return res, nil
}

func setSigningKey() (string, error) {
	rdb := GetRDB()

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	key := "asd"
	err = rdb.Set(ctx, "key", key, 0).Err()

	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing key:"+err.Error())
		fmt.Println("ERROR writing key:", err.Error())
		return "", err

	}

	producer.ProduceMessage("REDIS write", "Key set:"+key)
	return key, nil
}

func SetToken(token string, expireTime time.Duration) error {
	rdb := GetRDB()

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	err = rdb.Set(ctx, token, "0", expireTime*time.Second).Err()
	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing token:"+err.Error())
		return err
	}

	producer.ProduceMessage("REDIS write", "Token:"+token+" set")
	return nil
}

func GetToken(token string) (bool, error) {
	rdb := GetRDB()

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	_, err = rdb.Get(ctx, token).Result()
	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading token:"+token+", err: "+err.Error())
		return false, err
	}

	producer.ProduceMessage("REDIS write", "Token: "+token+" get")
	return true, nil
}
