package repository

import (
	"udonate/config"
	"udonate/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *UserRepository) GetAllUserConnections(userId string) (conList []entity.Connection, errorObject error) {
	return conList, nil
}

func (repository *UserRepository) GetConnectionById(connectionId string) (conObject entity.Connection, errorObject error) {
	return conObject, nil
}

func (repository *UserRepository) InsertNewConnection(connection entity.Connection) (newConnection entity.Connection, errorObject error) {
	return newConnection, nil
}

func (repository *UserRepository) UpdateUserConnections(connectionId string, connection entity.Connection) (updatedUser entity.User, errorObject error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	con := map[string]interface{}{}
	if connection.Title != "" {
		con["connections.0.title"] = connection.Title
	}
	if connection.Value != "" {
		con["connections.0.value"] = connection.Value
	}
	if connection.Meta != "" {
		con["connections.0.meta"] = connection.Meta
	}

	var updatedDocument bson.M
	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{"connections._id", connectionId}}
	updateQuery := map[string]interface{}{
		"$set": con,
	}
	err := repository.Collection.FindOneAndUpdate(ctx, filter,
		updateQuery,
		opts,
	).Decode(&updatedDocument)
	// exception.PanicIfNeeded(err)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
