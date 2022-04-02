package repository

import (
	"udonate/config"
	"udonate/entity"
	"udonate/exception"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) IUserRepository {
	return &UserRepository{
		Collection: database.Collection("Users"),
	}
}

type UserRepository struct {
	Collection *mongo.Collection
}

func (repository *UserRepository) FindAll() (Users []entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		user := entity.User{}
		if mapstructure.Decode(document, &user) != nil {
			exception.PanicIfNeeded(err)
		}
		Users = append(Users, user)
	}
	return Users
}

func (repository *UserRepository) FindUserById(userId string) (selectedUser entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{
		"_id": userId,
	})
	exception.PanicIfNeeded(err)

	var document bson.M
	err = cursor.All(ctx, &document)
	exception.PanicIfNeeded(err)

	if mapstructure.Decode(document, &selectedUser) != nil {
		exception.PanicIfNeeded(err)
	}
	return selectedUser
}

func (repository *UserRepository) Insert(User entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.InsertOne(ctx, User)
	exception.PanicIfNeeded(err)
}

func (repository *UserRepository) Delete(userId string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{
		"_id": userId,
	})
	exception.PanicIfNeeded(err)
}
