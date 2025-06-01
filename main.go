package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	expenseApi "github.com/IgorCooli/xpense/api/expense"
	expenseService "github.com/IgorCooli/xpense/internal/business/service/expense"
	expenseRepository "github.com/IgorCooli/xpense/internal/repository/expense"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbClient := setupDb(ctx)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
	}))

	injectExpenseApi(ctx, dbClient, app)

	port := resolveApiPort()

	app.Listen(":" + port)
}

func setupDb(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:mongodb159@tccmongodb.3ud5x.mongodb.net/?retryWrites=true&w=majority&appName=TCCMongoDB"))
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("Could not connect to database")
	}
	return client
}

func generateSecretKey() (string, error) {
	length := 32
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	secretKey := base64.URLEncoding.EncodeToString(randomBytes)
	return secretKey, nil
}

func injectExpenseApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {
	expenseRepository := expenseRepository.NewRepository(dbClient)
	expenseService := expenseService.NewService(expenseRepository)
	expenseApi.NewHandler(ctx, expenseService, app)
}

func resolveApiPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
