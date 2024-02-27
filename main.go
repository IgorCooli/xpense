package main

import (
	"context"
	"os"
	"time"

	cardApi "github.com/IgorCooli/xpense/api/card"
	expenseApi "github.com/IgorCooli/xpense/api/expense"
	userApi "github.com/IgorCooli/xpense/api/user"
	cardService "github.com/IgorCooli/xpense/internal/business/service/card"
	expenseService "github.com/IgorCooli/xpense/internal/business/service/expense"
	userService "github.com/IgorCooli/xpense/internal/business/service/user"
	cardRepository "github.com/IgorCooli/xpense/internal/repository/card"
	expenseRepository "github.com/IgorCooli/xpense/internal/repository/expense"
	userRepository "github.com/IgorCooli/xpense/internal/repository/user"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbClient := setupDb(ctx)

	app := fiber.New()

	injectExpenseApi(ctx, dbClient, app)
	injectUserApi(ctx, dbClient, app)
	injectCardApi(ctx, dbClient, app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}

func setupDb(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:mongodb159@tccmongodb.3ud5x.mongodb.net/?retryWrites=true&w=majority&appName=TCCMongoDB"))
	if err != nil {
		panic("Could not connect to dabase")
	}

	return client
}

func injectExpenseApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {
	expenseRepository := expenseRepository.NewRepository(dbClient)
	expenseService := expenseService.NewService(expenseRepository)
	expenseApi.NewHandler(ctx, expenseService, app)
}

func injectUserApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {
	userRepository := userRepository.NewRepository(dbClient)
	userService := userService.NewService(userRepository)
	userApi.NewHandler(ctx, userService, app)
}

func injectCardApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {
	cardRepository := cardRepository.NewCardRepository(dbClient)
	cardService := cardService.NewService(cardRepository)
	cardApi.NewHandler(ctx, cardService, app)
}
