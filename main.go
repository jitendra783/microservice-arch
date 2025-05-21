package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{

			"message": "Hello, this is my Go microservice!",
		})
	})
	routes := router.Group("/api")
	routes.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, this is my Go microservice!",
		})
	})

	router.Run(":8086")

}
