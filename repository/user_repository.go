package repository

import (
	"udonate/config"
	"udonate/entity"
	"udonate/exception"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		user.Id = document["_id"].(string)
		Users = append(Users, user)
	}
	return Users
}

func (repository *UserRepository) FindUserById(userId string) (selectedUser entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var document bson.M
	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})
	err := repository.Collection.FindOne(ctx, bson.M{
		"_id": userId,
	}, opts).Decode(&document)
	exception.PanicIfNeeded(err)

	if mapstructure.Decode(document, &selectedUser) != nil {
		exception.PanicIfNeeded(err)
	}
	selectedUser.Id = document["_id"].(string)

	return selectedUser
}

func (repository *UserRepository) Insert(User entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.InsertOne(ctx, User)
	exception.PanicIfNeeded(err)
}

func (repository *UserRepository) Update(userId string, User entity.User) (updatedUser entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var updatedDocument bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", userId}}
	updateQuery := bson.D{{"$set", bson.D{
		{"name", User.Name},
	}}}
	err := repository.Collection.FindOneAndUpdate(ctx, filter,
		updateQuery,
		opts,
	).Decode(&updatedDocument)
	exception.PanicIfNeeded(err)

	if mapstructure.Decode(updatedDocument, &updatedUser) != nil {
		exception.PanicIfNeeded(err)
	}
	updatedUser.Id = updatedDocument["_id"].(string)

	return updatedUser
}

func (repository *UserRepository) Delete(userId string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{
		"_id": userId,
	})
	exception.PanicIfNeeded(err)
}
