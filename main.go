package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Halo blue",
		})
	})

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Gagal memulai server")
	}
}
