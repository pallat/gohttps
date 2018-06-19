package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": "pallat",
			"name":     "pallat",
			"lastname": "anchaleechamaikorn",
		})
	})
	r.RunTLS(":443", "local.crt", "server.key") // listen and serve on 0.0.0.0:8080
}
