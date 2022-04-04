package service

import (
	"time"
	"udonate/entity"
	"udonate/pkg/gdpr"
	"udonate/pkg/random_number"
	"udonate/repository"
	"udonate/validation"
	"udonate/view_model"

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

func (service *UserService) List() (responses []view_model.GetUserResponse) {
	Users := service.UserRepository.FindAll()
	for _, user := range Users {
		temp := view_model.GetUserResponse{
			Id:           user.Id,
			Name:         user.Name,
			Family:       user.Family,
			Nationality:  user.Nationality,
			NationalCode: gdpr.ShowGDPR(),
			Username:     user.Username,
			Sex:          user.Sex,
			ReferralCode: user.ReferralCode,
			Connections:  user.Connections,
			Status:       user.Status,
			CreatedAt:    user.CreatedAt,
		}
		temp.NationalCode = gdpr.ShowGDPR()
		for i := 0; i < len(temp.Connections); i++ {
			temp.Connections[i].VerificationCode = gdpr.ShowGDPR()
		}
		responses = append(responses, temp)
	}
	return responses
}

func (service *UserService) FindUser(userId string) (response view_model.GetUserResponse) {
	User, _ := service.UserRepository.FindUserById(userId)
	response = view_model.GetUserResponse{
		Id:           User.Id,
		Name:         User.Name,
		Family:       User.Family,
		Nationality:  User.Nationality,
		NationalCode: User.NationalCode,
		Username:     User.Username,
		Sex:          User.Sex,
		ReferralCode: User.ReferralCode,
		Connections:  User.Connections,
		Status:       User.Status,
		CreatedAt:    User.CreatedAt,
	}
	User.NationalCode = gdpr.ShowGDPR()
	for i := 0; i < len(response.Connections); i++ {
		response.Connections[i].VerificationCode = gdpr.ShowGDPR()
	}
	return response
}

func (service *UserService) Create(request view_model.CreateUserRequest) (response view_model.CreateUserResponse) {
	request.Id = uuid.New().String()
	validation.InsertNewUserValidation(request)

	user := entity.User{
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
				Title:            "email",
				Value:            request.Email,
				VerificationCode: random_number.RandomNumber(8),
				IsVerified:       false,
			},
			{
				Title:            "mobile",
				Value:            request.Phone,
				VerificationCode: random_number.RandomNumber(8),
				IsVerified:       false,
			},
		},
		Status: "ACTIVE",
	}
	service.UserRepository.Insert(user)

	response = view_model.CreateUserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Family:       user.Family,
		Nationality:  user.Nationality,
		NationalCode: user.NationalCode,
		Birthday:     user.Birthday,
		Username:     user.Username,
		Password:     gdpr.ShowGDPR(),
		Sex:          user.Sex,
		ReferralCode: user.ReferralCode,
		Connections:  user.Connections,
		Status:       user.Status,
		CreatedAt:    user.CreatedAt,
	}
	for i := 0; i < len(response.Connections); i++ {
		response.Connections[i].VerificationCode = gdpr.ShowGDPR()
	}
	return response
}

func (service *UserService) Update(userId string, request view_model.UpdateUserRequest) (response view_model.CreateUserResponse) {
	validation.UpdateUserValidation(request)

	user := entity.User{
		Id:           uuid.New().String(),
		Name:         request.Name,
		Family:       request.Family,
		Nationality:  request.Nationality,
		NationalCode: request.NationalCode,
		Birthday:     request.Birthday,
		Sex:          request.Sex,
		ReferralCode: request.ReferralCode,
		ModifiedAt:   time.Now().Format(time.RFC3339),
	}
	updatedUser, _ := service.UserRepository.Update(userId, user)

	response = view_model.CreateUserResponse{
		Id:           updatedUser.Id,
		Name:         updatedUser.Name,
		Family:       updatedUser.Family,
		Nationality:  updatedUser.Nationality,
		NationalCode: updatedUser.NationalCode,
		Birthday:     updatedUser.Birthday,
		Username:     updatedUser.Username,
		Password:     updatedUser.Password,
		Sex:          updatedUser.Sex,
		ReferralCode: updatedUser.ReferralCode,
		Connections:  updatedUser.Connections,
		Status:       updatedUser.Status,
		CreatedAt:    updatedUser.CreatedAt,
	}
	return response
}

func (service *UserService) Delete(userId string) (response view_model.GetUserResponse) {
	user, _ := service.UserRepository.DeleteUserById(userId)
	response = view_model.GetUserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Family:       user.Family,
		Nationality:  user.Nationality,
		NationalCode: user.NationalCode,
		Username:     user.Username,
		Connections:  user.Connections,
		Requests:     user.Requests,
		Sex:          user.Sex,
	}
	user.NationalCode = gdpr.ShowGDPR()
	return response
}

func (service *UserService) UpdateConnection(conId string, connection entity.Connection) (response view_model.GetUserResponse) {
	updatedUser, _ := service.UserRepository.UpdateUserConnections(conId, connection)
	response = view_model.GetUserResponse{
		Id:           updatedUser.Id,
		Name:         updatedUser.Name,
		Family:       updatedUser.Family,
		Nationality:  updatedUser.Nationality,
		NationalCode: updatedUser.NationalCode,
		Username:     updatedUser.Username,
		Connections:  updatedUser.Connections,
		Requests:     updatedUser.Requests,
		Sex:          updatedUser.Sex,
	}
	return response
}
