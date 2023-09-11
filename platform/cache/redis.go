package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"lido-core/v1/platform/database"

	"github.com/redis/go-redis/v9"
)

// RedisConnection func for connect to Redis server.

var client *redis.ClusterClient

func init() {
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			os.Getenv("REDIS_NODE_1"),
			os.Getenv("REDIS_NODE_2"),
			os.Getenv("REDIS_NODE_3"),
		},
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	log.Println("Connected to Redis Cluster:", pong)
	log.Println("[REDIS] Cluster connection successful")
}

func Shutdown() {
	if client != nil {
		// ctx, _ := database.NewContext()
		client.Close()
		log.Println("[CACHE] Shutdown client connection")
	}
}

func Set(key string, data string, expire time.Duration) error {
	ctx, _ := database.NewContext()
	return client.Set(ctx, key, data, expire).Err()
}

func Get(key string) (string, error) {
	ctx, _ := database.NewContext()
	return client.Get(ctx, key).Result()
}

func GetYoutubeView(key string) (uint64, error) {
	ctx, _ := database.NewContext()
	return client.Get(ctx, key).Uint64()
}

func Save(key string, data string) error {
	ctx, _ := database.NewContext()
	return client.Set(ctx, key, data, 0).Err()
}

func IsUnique(code string) bool {
	ctx, _ := database.NewContext()
	exists, err := client.Exists(ctx, code).Result()
	if err != nil {
		panic(err)
	}
	return exists == 0
}

func Incr(key string) (int64, error) {
	ctx, _ := database.NewContext()
	return client.Incr(ctx, key).Result()
}
