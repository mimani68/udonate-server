package repository

import (
	"errors"
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

func (repository *UserRepository) FindAll() (users []entity.User) {
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
		users = append(users, user)
	}
	return users
}

func (repository *UserRepository) FindUserById(userId string) (selectedUser entity.User, errorObject error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var document bson.M
	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})
	err := repository.Collection.FindOne(ctx, bson.M{
		"_id": userId,
	}, opts).Decode(&document)
	if err != nil {
		// exception.PanicIfNeeded(err)
		return selectedUser, err
	}

	if mapstructure.Decode(document, &selectedUser) != nil {
		// exception.PanicIfNeeded(err)
		return selectedUser, err
	}
	selectedUser.Id = document["_id"].(string)

	return selectedUser, nil
}

func (repository *UserRepository) Insert(user entity.User) (newUser entity.User, errorObject error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.InsertOne(ctx, user)
	// exception.PanicIfNeeded(err)
	if err != nil {
		return newUser, err
	} else {
		return user, nil
	}
}

func (repository *UserRepository) Update(userId string, User entity.User) (updatedUser entity.User, errorObject error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	update := map[string]string{}
	if User.Name != "" {
		update["name"] = User.Name
	}
	if User.Family != "" {
		update["family"] = User.Family
	}
	if User.Sex != "" {
		update["sex"] = User.Sex
	}
	if User.Nationality != "" {
		update["nationality"] = User.Nationality
	}
	if User.NationalCode != "" {
		update["nationalCode"] = User.NationalCode
	}
	if User.Birthday != "" {
		update["birthday"] = User.Birthday
	}
	if User.ReferralCode != "" {
		update["referralCode"] = User.ReferralCode
	}
	if len(update) <= 0 {
		return updatedUser, errors.New("There is no new field for update table user.")
	}

	var updatedDocument bson.M
	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{"_id", userId}}
	updateQuery := map[string]interface{}{
		"$set": update,
	}
	err := repository.Collection.FindOneAndUpdate(ctx, filter,
		updateQuery,
		opts,
	).Decode(&updatedDocument)
	// exception.PanicIfNeeded(err)
	if err != nil {
		return updatedUser, err
	}

	errFind := repository.Collection.FindOne(ctx, map[string]string{
		"_id": userId,
	}).Decode(&updatedDocument)
	// exception.PanicIfNeeded(errFind)
	if errFind != nil {
		return updatedUser, errFind
	}

	if mapstructure.Decode(updatedDocument, &updatedUser) != nil {
		// exception.PanicIfNeeded(err)
		return updatedUser, errFind
	}
	updatedUser.Id = updatedDocument["_id"].(string)

	return updatedUser, nil
}

func (repository *UserRepository) DeleteUserById(userId string) (deletedUser entity.User, errorObject error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{
		"_id": userId,
	})
	if err != nil {
		return deletedUser, err
	} else {
		return deletedUser, nil
	}
}
