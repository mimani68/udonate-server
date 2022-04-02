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
		connectionStructList := []entity.Connection{}
		if mapstructure.Decode(document["connections"], &connectionStructList) != nil {
			exception.PanicIfNeeded(err)
		}
		reuqestStructList := []entity.Request{}
		if mapstructure.Decode(document["requests"], &reuqestStructList) != nil {
			exception.PanicIfNeeded(err)
		}
		Users = append(Users, entity.User{
			Id:                 document["_id"].(string),
			Name:               document["name"].(string),
			Family:             document["family"].(string),
			Nationality:        document["nationality"].(string),
			Birthday:           document["birthday"].(string),
			Sex:                document["sex"].(string),
			NationalCode:       document["nationalCode"].(string),
			Password:           document["password"].(string),
			Connections:        connectionStructList,
			ReferralCode:       document["referralCode"].(string),
			Requests:           reuqestStructList,
			Status:             document["status"].(string),
			CreatedAt:          document["createdAt"].(string),
			ModifiedAt:         document["modifiedAt"].(string),
			DeletedAt:          document["deletedAt"].(string),
			DeletedDescription: document["deletedDescription"].(string),
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
