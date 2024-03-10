package main

import (
	"github.com/gin-gonic/gin"
	"usermgtms/controller/UserController"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/usermgt/user", UserController.UserCreate)

	router.Run()
}
