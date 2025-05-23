package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init findUserByID repository.",
		zap.String("journey", "findUserByID"),
	)

	collection_name := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)

	collection := ur.dataBaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s.", id)
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByID"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID."
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByID"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByID repository executed suceeefully.",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init findUserByEmail repository.",
		zap.String("journey", "findUserByEmail"),
	)

	collection_name := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)

	collection := ur.dataBaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s.", email)
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByEmail"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email."
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed suceeefully.",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
