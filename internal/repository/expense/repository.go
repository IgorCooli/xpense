package expense

import (
	"context"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOne(ctx context.Context, expense model.Expense) error
}

func NewRepository(client *mongo.Client) Repository {
	return mongoRepository{
		expenseDB: client.Database("TCCMongoDB").Collection("expense"),
	}
}

type mongoRepository struct {
	expenseDB *mongo.Collection
}

func (r mongoRepository) InsertOne(ctx context.Context, expense model.Expense) error {

	result, err := r.expenseDB.InsertOne(ctx, expense)

	if err != nil {
		panic("Could not insert item")
	}

	fmt.Println(result)
	return nil
}
