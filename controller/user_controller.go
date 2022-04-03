package controller

import (
	"udonate/exception"
	"udonate/service"
	"udonate/view_model"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(UserService *service.IUserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) Route(app fiber.Router) {
	app.Post("/users", controller.Create)
	app.Get("/me", controller.Me)
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
	var request view_model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response := controller.UserService.Create(request)
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Me(c *fiber.Ctx) error {
	userId := c.Get("USER")
	responses := controller.UserService.FindUser(userId)
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserController) List(c *fiber.Ctx) error {
	responses := controller.UserService.List()
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
