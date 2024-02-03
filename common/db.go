package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func GetDbConnection() *gorm.DB {
	var err error
	err = godotenv.Load()
	if err != nil {
		return nil
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	database, _ := gorm.Open(postgres.Open(url), &gorm.Config{})
	return database
}

func GetRedisConnection() *redis.Client {
	var err error
	err = godotenv.Load()
	if err != nil {
		return nil
	}
	redisAdd := os.Getenv("REDIS_ADDRESS")
	redisUser := os.Getenv("REDIS_USER")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 0, 8)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAdd,
		Username: redisUser,
		Password: redisPassword,
		DB:       int(redisDB),
	})

	return redisClient
}
