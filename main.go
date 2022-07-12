package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.SetTrustedProxies(nil)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
