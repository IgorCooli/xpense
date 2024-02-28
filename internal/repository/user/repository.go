package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Respository interface {
	InsertOne(ctx context.Context, user model.User) error
	FindByUsername(ctx context.Context, userName string) (model.User, error)
}

type mongoRepository struct {
	userDB *mongo.Collection
}

func NewRepository(client *mongo.Client) Respository {
	return mongoRepository{
		userDB: client.Database("TCCMongoDB").Collection("user"),
	}
}

func (r mongoRepository) FindByUsername(ctx context.Context, userName string) (model.User, error) {
	var result model.User
	filter := bson.D{
		{"username", userName},
	}

	r.userDB.FindOne(ctx, filter).Decode(&result)

	if result.ID == "" {
		return model.User{}, errors.New("User not found")
	}

	return result, nil
}

func (r mongoRepository) InsertOne(ctx context.Context, user model.User) error {

	result, err := r.userDB.InsertOne(ctx, user)

	if err != nil {
		panic("Could not insert item")
	}

	fmt.Println(result)
	return nil
}
