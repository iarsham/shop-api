package db

import (
	"context"
	"fmt"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/redis/go-redis/v9"
	"os"
)

var (
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	Client    *redis.Client
)

func RedisClient() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", RedisHost, RedisPort),
		Password: "",
		DB:       0,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return Client
}

func CloseRedis(logs *common.Logger) {
	err := Client.Close()
	logs.Warn(err.Error())
}
