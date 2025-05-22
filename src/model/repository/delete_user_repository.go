package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info(
		"Init deleteUser repository.",
		zap.String("journey", "deleteUser"),
	)

	collection_name := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	collection := ur.dataBaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error(
			"Error trying to delete user.",
			err,
			zap.String("journey", "deleteUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"deleteUser repository executed suceeefully.",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
