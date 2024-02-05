package middlewares

import (
	"github.com/Bluhabit/uwang-rest-storage/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := authHeader{}

		if err := context.ShouldBindHeader(&header); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				err := "Invalid request parameter"
				context.JSON(401, gin.H{
					"status_code": 401,
					"data":        nil,
					"message":     err,
				})
				context.Abort()
				return
			}

			//error type unknown
			context.JSON(401, gin.H{
				"status_code": 401,
				"data":        nil,
				"message":     "500 Error",
			})
			context.Abort()
			return
		}

		idTokenHeader := strings.Split(header.Token, "Bearer ")

		if len(idTokenHeader) < 2 {
			context.JSON(401, gin.H{
				"status_code": 401,
				"data":        nil,
				"message":     "Token not provided",
			})
			context.Abort()
			return
		}

		//validate token
		claims := common.DecodeJWT(idTokenHeader[1])
		if claims == nil {
			context.JSON(401, gin.H{
				"status_code": 401,
				"data":        nil,
				"message":     "Token not valid",
			})
			context.Abort()
			return
		}
		context.Set("session_id", claims.Sub)
		context.Next()
	}
}
