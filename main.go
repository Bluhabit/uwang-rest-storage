package main

import (
	"github.com/Bluhabit/uwang-rest-storage/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Halo blue",
		})
	})
	
	router.Run(":8080")

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Gagal memulai server")
	}
}
