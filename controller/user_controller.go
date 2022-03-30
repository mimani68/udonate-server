package controller

import (
	"udonate/exception"
	"udonate/model"
	"udonate/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(UserService *service.IUserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/api/users", controller.Create)
	app.Get("/api/me", controller.List)
}

func (controller *UserController) ConsoleRoute(app *fiber.App) {
	app.Get("/console/users", controller.List)
	app.Get("/console/users/:userId", controller.List)
	app.Post("/console/users", controller.Create)
	app.Put("/console/users/:userId", controller.List)
	app.Patch("/console/users/:userId", controller.List)
	app.Delete("/console/users/:userId", controller.List)
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.UserService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) List(c *fiber.Ctx) error {
	responses := controller.UserService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
