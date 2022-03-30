package service

import "udonate/model"

type IUserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
	List() (responses []model.GetUserResponse)
}
