package main

import (
	"log"
	"net/http"

	"github.com/Bluhabit/uwang-rest-storage/common"
	"github.com/Bluhabit/uwang-rest-storage/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	common.GetDbConnection()
	common.GetRedisConnection()
	common.GetMinioClient()
	routes.InitRoutes(router)

	router.GET("/", func(c *gin.Context) {
		dec := common.DecodeJWT("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJibHVoYWJpdC5pZCIsInN1YiI6IlRyaWFuIiwiaWF0IjoxNzA3MDk3ODM4LCJleHAiOjE3MDcxMDE0Mzh9.XFB7fjyeHffSrypl4NRO63J-RxZHgA-fRmt7ZIsxRIQ")
		c.JSON(http.StatusOK, gin.H{
			"Message": dec.Sub,
		})
	})

	router.Run(":8080")

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Gagal memulai server")
	}
}
