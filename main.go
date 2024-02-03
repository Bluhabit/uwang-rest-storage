package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Halo blue",
		})
	})
	router.POST("/post/:name", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s, message: %s", id, page, name, message)
		context.JSON(200,gin.H{
			"status_code": 1001, //status kode merepresentasikan kondisi
			"data": true,
			"message": "Data Diterima",
		})
	})
	router.Run(":8080")

	router.POST("/upload-avatar", func(context *gin.Context) {

	})
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Gagal memulai server")
	}
}
