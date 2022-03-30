package controller

import (
	"udonate/config"
	"udonate/repository"
	"udonate/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	userController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMongoDatabase(configuration)
var UserRepository = repository.NewUserRepository(database)
var UserService = service.NewUserService(&UserRepository)

var userController = NewUserController(&UserService)

var app = createTestApp()
