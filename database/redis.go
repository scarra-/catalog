package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func NewRedis() *Redis {
	client := NewRedisConnection(0)

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	r := &Redis{Client: client}

	return r
}

func NewPrometheusRedis() *redis.Client {
	client := NewRedisConnection(1)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	return client
}

func NewRedisConnection(dbId int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbId,
	})
}
