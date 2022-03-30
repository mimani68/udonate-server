package service

import (
	"udonate/entity"
	"udonate/model"
	"udonate/repository"
	"udonate/validation"
)

func NewUserService(UserRepository *repository.IUserRepository) IUserService {
	return &UserServiceImpl{
		UserRepository: *UserRepository,
	}
}

type UserServiceImpl struct {
	UserRepository repository.IUserRepository
}

func (service *UserServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {
	validation.Validate(request)

	User := entity.User{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.UserRepository.Insert(User)

	response = model.CreateUserResponse{
		Id:       User.Id,
		Name:     User.Name,
		Price:    User.Price,
		Quantity: User.Quantity,
	}
	return response
}

func (service *UserServiceImpl) List() (responses []model.GetUserResponse) {
	Users := service.UserRepository.FindAll()
	for _, User := range Users {
		responses = append(responses, model.GetUserResponse{
			Id:       User.Id,
			Name:     User.Name,
			Price:    User.Price,
			Quantity: User.Quantity,
		})
	}
	return responses
}
