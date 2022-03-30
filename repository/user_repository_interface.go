package repository

import "udonate/entity"

type IUserRepository interface {
	Insert(User entity.User)

	FindAll() (Users []entity.User)

	DeleteAll()
}
