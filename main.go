package main

import (
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

	// Setup Routing
	UserController.Route(app)

	// Start App
	err := app.Listen(configuration.Get("ADDRESS") + ":" + configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
