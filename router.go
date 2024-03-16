package main

import (
	"go-auth-server/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func router(app *fiber.App) {
	app.Get("/", Hello)
	app.Get("/users", controllers.GetAllUsers)
	// e.GET("/validation", middlewares.RequireAuth, controllers.Validate)

	// e.POST("/signup", controllers.SignUp)
	// e.POST("/login", controllers.Login)
}

func Hello(c *fiber.Ctx) error {
	c.SendStatus(http.StatusOK)
	c.JSON(fiber.Map{
		"message": "Hello World",
	})

	return nil
}
