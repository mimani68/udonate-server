package service

import model "udonate/view_model"

type IUserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
	List() (responses []model.GetUserResponse)
}
