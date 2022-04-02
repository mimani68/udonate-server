package main

import (
	"time"
	"udonate/config"
	"udonate/controller"
	"udonate/exception"
	"udonate/repository"
	"udonate/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New("./.env")
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	UserRepository := repository.NewUserRepository(database)

	// Setup Service
	UserService := service.NewUserService(&UserRepository)

	// Setup Controller
	UserController := controller.NewUserController(&UserService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	apiRoute := app.Group("/v1")

	// Ping
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "PONG",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// Setup Routing
	UserController.Route(apiRoute)
	UserController.ConsoleRoute(app)

	// Start App
	// err := app.Listen("0.0.0.0:" + configuration.Get("PORT"))
	err := app.Listen(":" + configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
