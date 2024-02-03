package common

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
)

type UserClaims struct {
	Id  string `json:"id"`
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
	jwt.Claims
}

func EncodeJWT(claims UserClaims) string {
	var err error
	err = godotenv.Load()
	if err != nil {
		return ""
	}
	key := []byte(os.Getenv("JWT_SECRET"))

	encoder := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	encoded, _ := encoder.SignedString(key)
	return encoded
}

func DecodeJWT(token string) *UserClaims {
	var err error
	err = godotenv.Load()
	if err != nil {
		return nil
	}
	key := []byte(os.Getenv("JWT_SECRET"))
	decoder, _ := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return decoder.Claims.(*UserClaims)

}
