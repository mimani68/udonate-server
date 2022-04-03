package main

import (
	"time"
	"udonate/config"
	"udonate/controller"
	"udonate/exception"
	"udonate/middleware"
	"udonate/repository"
	"udonate/service"

	"github.com/gofiber/fiber/v2"
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

	// Middleware
	middleware.GeneralMiddleWare(app, configuration)
	middleware.Auth(app)

	// Prefix of app
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
	err := app.Listen(configuration.Get("ADDRESS") + ":" + configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
