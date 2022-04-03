package service

import model "udonate/view_model"

type IUserService interface {
	List() (responses []model.GetUserResponse)
	FindUser(userId string) (response model.GetUserResponse)
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
	Update(userId string, request model.UpdateUserRequest) (response model.CreateUserResponse)
	Delete(userId string) (response model.GetUserResponse)
}
