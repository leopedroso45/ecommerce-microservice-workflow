package db

import (
	"github.com/go-redis/redis/v8"
	"os"
)

//set REDIS_HOST=localhost:6379
func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:               os.Getenv("REDIS_HOST"),
		Password:           "",
		DB:                 0,
	})
	return client
}
