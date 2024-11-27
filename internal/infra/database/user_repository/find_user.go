package user_repository

import (
	"context"
	"errors"

	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/entity/user_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) user_entity.UserRepositoryInterface {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {
	var userMongo UserEntityMongo

	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(ctx, filter).Decode(&userMongo)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("User not found with ID: "+id, err)
			return nil, internal_error.NewNotFoundError("User not found with ID: " + id)
		}

		logger.Error("Error finding user with ID: "+id, err)
		return nil, internal_error.NewInternalServerError("Error finding user with ID: " + id)
	}

	userEntity := &user_entity.User{
		ID:   userMongo.ID,
		Name: userMongo.Name,
	}

	return userEntity, nil
}
