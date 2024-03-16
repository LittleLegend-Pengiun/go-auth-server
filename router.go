package main

import (
	"go-auth-server/controllers"
	"go-auth-server/middlewares"

	"github.com/gin-gonic/gin"
)

func router(e *gin.Engine) *gin.Engine {
	e.GET("/", Hello)
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/validation", middlewares.RequireAuth, controllers.Validate)

	e.POST("/signup", controllers.SignUp)
	e.POST("/login", controllers.Login)

	return e
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
