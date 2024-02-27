package user

import (
	"context"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Respository interface {
	InsertOne(ctx context.Context, user model.User) error
}

type mongoRepository struct {
	userDB *mongo.Collection
}

func NewUserRepository(client *mongo.Client) Respository {
	return mongoRepository{
		userDB: client.Database("TCCMongoDB").Collection("user"),
	}
}

func (r mongoRepository) InsertOne(ctx context.Context, user model.User) error {

	result, err := r.userDB.InsertOne(ctx, user)

	if err != nil {
		panic("Could not insert item")
	}

	fmt.Println(result)
	return nil
}
