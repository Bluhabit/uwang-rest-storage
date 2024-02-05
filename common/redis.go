package common

import (
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var (
	cache *redis.Client
)

func GetRedisConnection() *redis.Client {
	if cache == nil {
		var err error
		err = godotenv.Load()
		if err != nil {
			return nil
		}
		redisAdd := os.Getenv("REDIS_ADDRESS")
		redisUser := os.Getenv("REDIS_USER")
		redisPassword := os.Getenv("REDIS_PASSWORD")
		redisDB, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 0, 8)

		cache = redis.NewClient(&redis.Options{
			Addr:     redisAdd,
			Username: redisUser,
			Password: redisPassword,
			DB:       int(redisDB),
		})
	}
	return cache
}
