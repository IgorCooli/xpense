package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	authApi "github.com/IgorCooli/xpense/api/auth"
	cardApi "github.com/IgorCooli/xpense/api/card"
	expenseApi "github.com/IgorCooli/xpense/api/expense"
	userApi "github.com/IgorCooli/xpense/api/user"
	cardService "github.com/IgorCooli/xpense/internal/business/service/card"
	expenseService "github.com/IgorCooli/xpense/internal/business/service/expense"
	"github.com/IgorCooli/xpense/internal/business/service/helpers/jwt"
	passwordService "github.com/IgorCooli/xpense/internal/business/service/helpers/password"
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

	jwtSecret, err := generateSecretKey()

	if err != nil {
		panic("Error generating secret key")
	}

	injectExpenseApi(ctx, dbClient, app)
	injectUserApi(ctx, dbClient, app, jwtSecret)
	injectCardApi(ctx, dbClient, app)
	injectAuthApi(ctx, dbClient, app, jwtSecret)

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

func injectUserApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App, jwtSecret string) {
	userRepository := userRepository.NewRepository(dbClient)
	passwordService := passwordService.NewPasswordService()

	jwtService := jwt.NewJwtService(jwtSecret)
	userService := userService.NewService(userRepository, passwordService, jwtService)
	userApi.NewHandler(ctx, userService, app)
}

func injectCardApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {
	cardRepository := cardRepository.NewCardRepository(dbClient)
	cardService := cardService.NewService(cardRepository)
	cardApi.NewHandler(ctx, cardService, app)
}

func injectAuthApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App, jwtSecret string) {
	userRepository := userRepository.NewRepository(dbClient)
	passwordService := passwordService.NewPasswordService()

	jwtService := jwt.NewJwtService(jwtSecret)
	userService := userService.NewService(userRepository, passwordService, jwtService)
	authApi.NewHandler(ctx, userService, app)
}
