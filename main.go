package main

import (
	"context"
	"os"
	"time"

	api "github.com/IgorCooli/xpense/api/expense"
	service "github.com/IgorCooli/xpense/internal/business/service/expense"
	repository "github.com/IgorCooli/xpense/internal/repository/expense"
	"github.com/gofiber/fiber/v3"
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

	repo := repository.NewRepository(dbClient)
	service := service.NewService(repo)

	app := fiber.New()

	api.NewHandler(ctx, service, app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Porta padrão se não estiver definida
	}

	app.Listen(":" + port)
}
