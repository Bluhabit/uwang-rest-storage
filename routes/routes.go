package routes

import (
	"github.com/Bluhabit/uwang-rest-storage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/upload-profile-picture", middlewares.AuthMiddleware(),UploadProfilePicture)
	}
}
