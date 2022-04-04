package service

import (
	"udonate/entity"
	"udonate/view_model"
)

type IUserService interface {
	List() (responses []view_model.GetUserResponse)
	FindUser(userId string) (response view_model.GetUserResponse)
	Create(request view_model.CreateUserRequest) (response view_model.CreateUserResponse)
	Update(userId string, request view_model.UpdateUserRequest) (response view_model.CreateUserResponse)
	Delete(userId string) (response view_model.GetUserResponse)

	// RequstDonate(userId string, request entity.Request) (response view_model.GetUserDonateReqResponse)

	UpdateConnection(conId string, connection entity.Connection) (response view_model.GetUserResponse)
}
