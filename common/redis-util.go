package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func CreateRedisKeyUserSession(SessionId string) string {
	var err error
	err = godotenv.Load()
	if err != nil {
		return ""
	}
	environment := os.Getenv("ENVIRONMENT")
	return fmt.Sprintf("%s:sign-in:%s", environment, SessionId)
}
