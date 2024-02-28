package expense

import (
	"context"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOne(ctx context.Context, expense model.Expense) error
	InsertMany(ctx context.Context, expenses []model.Expense) error
	Search(ctx context.Context, userId string) []model.Expense
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

func (r mongoRepository) InsertMany(ctx context.Context, expenses []model.Expense) error {

	var input []interface{}
	for _, exp := range expenses {
		input = append(input, exp)
	}

	result, err := r.expenseDB.InsertMany(ctx, input)
	if err != nil {
		panic("Could not insert items")
	}

	fmt.Println(result)
	return nil
}

func (r mongoRepository) Search(ctx context.Context, userId string) []model.Expense {
	var results []model.Expense
	filter := bson.D{
		{"card.userid", userId},
	}

	cursor, err := r.expenseDB.Find(ctx, filter)
	if err != nil {
		return []model.Expense{}
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result model.Expense
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		results = append(results, result)
	}

	return results
}
