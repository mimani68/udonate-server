package repository

import (
	"udonate/config"
	"udonate/entity"
	"udonate/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) IUserRepository {
	return &UserRepositoryImpl{
		Collection: database.Collection("Users"),
	}
}

type UserRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *UserRepositoryImpl) Insert(User entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":      User.Id,
		"name":     User.Name,
		"price":    User.Price,
		"quantity": User.Quantity,
	})
	exception.PanicIfNeeded(err)
}

func (repository *UserRepositoryImpl) FindAll() (Users []entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		Users = append(Users, entity.User{
			Id:       document["_id"].(string),
			Name:     document["name"].(string),
			Price:    document["price"].(int64),
			Quantity: document["quantity"].(int32),
		})
	}

	return Users
}

func (repository *UserRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}
