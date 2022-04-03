package controller

import (
	"fmt"
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
	app.Post("/users", controller.CreateUser)
	app.Get("/me", controller.Me)
}

func (controller *UserController) ConsoleRoute(app *fiber.App) {
	app.Get("/console/users", controller.UsersList)
	app.Get("/console/users/:userId", controller.FindUser)
	app.Post("/console/users", controller.CreateUser)
	app.Patch("/console/users/:userId", controller.UpdateUser)
	// app.Patch("/console/users/:userId", controller.List)
	// app.Delete("/console/users/:userId", controller.List)
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
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
	// userId := c.Get("USER")
	userId := "cba2a781-2536-4543-a22e-bfe3a0e3fd8c"
	if len(userId) < 10 {
		return c.JSON(view_model.WebResponse{
			Code:   400,
			Status: "NOK",
			Data:   nil,
		})
	}
	responses := controller.UserService.FindUser(userId)
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserController) UsersList(c *fiber.Ctx) error {
	responses := controller.UserService.List()
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserController) FindUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if len(userId) < 10 {
		return c.JSON(view_model.WebResponse{
			Code:   400,
			Status: "NOK",
			Data:   nil,
		})
	}
	responses := controller.UserService.FindUser(userId)
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserController) UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if len(userId) < 10 {
		return c.JSON(view_model.WebResponse{
			Code:   400,
			Status: "NOK",
			Data:   nil,
		})
	}
	user := view_model.UpdateUserRequest{}
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.JSON(view_model.WebResponse{
			Code:   400,
			Status: "NOK",
			Data:   err,
		})
	}
	if len(userId) < 10 {
		return c.JSON(view_model.WebResponse{
			Code:   400,
			Status: "NOK",
			Data:   nil,
		})
	}
	responses := controller.UserService.Update(userId, user)
	return c.JSON(view_model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
