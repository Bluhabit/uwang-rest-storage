package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDbConnection() *gorm.DB {
	var err error
	if len(os.Getenv("DB_HOST")) < 1 {
		err = godotenv.Load()
		if err != nil {
			fmt.Println(err)
			return nil
		}
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

	if db == nil {
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			log.Println("Failed connect to database")
			return nil
		}
		log.Println("Successfully connected to database")
	}
	return db
}
