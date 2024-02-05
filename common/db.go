package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDbConnection() *gorm.DB {
	if db == nil {
		var err error
		err = godotenv.Load()
		if err != nil {

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

		once.Do(func() {
			db, _ = gorm.Open(postgres.Open(url), &gorm.Config{})
		})
	}
	return db
}
