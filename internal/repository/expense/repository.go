package expense

import (
	"context"

	"github.com/IgorCooli/xpense/internal/business/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOne(ctx context.Context, expense model.Expense) error
	InsertMany(ctx context.Context, expenses []model.Expense) error
	Search(ctx context.Context, month string, year string) []model.Expense
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

	_, err := r.expenseDB.InsertOne(ctx, expense)

	if err != nil {
		panic("Could not insert item")
	}

	return nil
}

func (r mongoRepository) InsertMany(ctx context.Context, expenses []model.Expense) error {

	var input []interface{}
	for _, exp := range expenses {
		input = append(input, exp)
	}

	_, err := r.expenseDB.InsertMany(ctx, input)
	if err != nil {
		panic("Could not insert items")
	}

	return nil
}

func (r mongoRepository) Search(ctx context.Context, month string, year string) []model.Expense {
	var results []model.Expense

	filter := bson.D{
		{"month", month},
		{"year", year},
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
