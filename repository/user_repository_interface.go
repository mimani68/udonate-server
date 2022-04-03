package repository

import (
	"udonate/entity"
)

type IUserRepository interface {
	FindAll() (Users []entity.User)
	FindUserById(userId string) (selectedUser entity.User)
	Insert(User entity.User)
	Update(userId string, User entity.User) (updatedUser entity.User)
	// SoftDelete(userId string)
	Delete(userId string)
}
