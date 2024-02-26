package main

import (
	"context"
	"time"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/repository/expense"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:mongodb159@tccmongodb.3ud5x.mongodb.net/?retryWrites=true&w=majority&appName=TCCMongoDB"))
	if err != nil {
		panic("Could not connect to dabase")
	}

	repo := expense.NewRepository(dbClient)

	model := model.Expense{
		ID:           "1",
		Value:        10.00,
		PaymentDate:  time.Now(),
		Installments: 1,
	}

	repo.InsertOne(ctx, model)
}
