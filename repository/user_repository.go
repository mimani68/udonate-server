package repository

import (
	"udonate/config"
	"udonate/entity"
	"udonate/exception"

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

func (repository *UserRepository) Insert(User entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.InsertOne(ctx, User)
	exception.PanicIfNeeded(err)
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
		Users = append(Users, entity.User{
			Id:   document["_id"].(string),
			Name: document["name"].(string),
		})
	}

	return Users
}

func (repository *UserRepository) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}
