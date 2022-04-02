package service

import (
	"time"
	"udonate/entity"
	"udonate/repository"
	"udonate/validation"
	model "udonate/view_model"

	"github.com/google/uuid"
)

func NewUserService(UserRepository *repository.IUserRepository) IUserService {
	return &UserService{
		UserRepository: *UserRepository,
	}
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func (service *UserService) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {
	request.Id = uuid.New().String()
	validation.Validate(request)

	User := entity.User{
		Id:           uuid.New().String(),
		Name:         request.Name,
		Family:       request.Family,
		Nationality:  request.Nationality,
		NationalCode: request.NationalCode,
		Birthday:     request.Birthday,
		Username:     request.Username,
		Password:     request.Password,
		Sex:          request.Sex,
		ReferralCode: request.ReferralCode,
		CreatedAt:    time.Now().Format(time.RFC3339),
		Connections: []entity.Connection{
			{
				Title:      "email",
				Value:      request.Email,
				IsVerified: false,
			},
			{
				Title:      "mobile",
				Value:      request.Phone,
				IsVerified: false,
			},
		},
		Status: "ACTIVE",
	}
	service.UserRepository.Insert(User)

	response = model.CreateUserResponse{
		Id:           User.Id,
		Name:         User.Name,
		Family:       User.Family,
		Nationality:  User.Nationality,
		NationalCode: User.NationalCode,
		Birthday:     User.Birthday,
		Username:     User.Username,
		Password:     User.Password,
		Sex:          User.Sex,
		ReferralCode: User.ReferralCode,
		Connections:  User.Connections,
		Status:       User.Status,
	}
	return response
}

func (service *UserService) List() (responses []model.GetUserResponse) {
	Users := service.UserRepository.FindAll()
	for _, user := range Users {
		temp := model.GetUserResponse{
			Id:           user.Id,
			Name:         user.Name,
			Family:       user.Family,
			Nationality:  user.Nationality,
			NationalCode: user.NationalCode,
			Username:     user.Username,
			Connections:  user.Connections,
			Sex:          user.Sex,
		}
		temp.NationalCode = "***"
		responses = append(responses, temp)
	}
	return responses
}
