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
	app.Post("/api/Users", controller.Create)
	app.Get("/api/Users", controller.List)
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
