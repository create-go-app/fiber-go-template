package cache

import (
	"os"
	"strconv"

	"github.com/create-go-app/fiber-go-template/pkg/utils"

	"github.com/go-redis/redis/v8"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() (*redis.Client, error) {
	// Define Redis database number.
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}

	// Set Redis options.
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	return redis.NewClient(options), nil
}
