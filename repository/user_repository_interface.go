package repository

import (
	"udonate/entity"
)

type IUserRepository interface {
	FindAll() (user []entity.User)
	FindUserById(userId string) (selectedUser entity.User, errorObject error)
	Insert(user entity.User) (newUser entity.User, errorObject error)
	Update(userId string, user entity.User) (updatedUser entity.User, errorObject error)
	// SoftDelete(userId string)
	DeleteUserById(userId string) (deletedUser entity.User, errorObject error)
}
