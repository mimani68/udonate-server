package controller

import (
	"udonate/exception"
	"udonate/model"
	"udonate/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserConsoleController struct {
	UserService service.IUserService
}

func NewUserConsoleController(UserService *service.IUserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserConsoleController) Route(app *fiber.App) {
	app.Post("/console/users", controller.Create)
	// app.Get("/console/user/:id", controller.List)
	app.Get("/console/users", controller.List)
	// app.Delete("/console/users", controller.List)
	// app.Patch("/console/users", controller.List)
	// app.Put("/console/users", controller.List)
}

func (controller *UserConsoleController) Create(c *fiber.Ctx) error {
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

func (controller *UserConsoleController) List(c *fiber.Ctx) error {
	responses := controller.UserService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
