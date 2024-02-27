package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	service "github.com/IgorCooli/xpense/internal/business/service/expense"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service service.Service
}

func NewHandler(ctx context.Context, service service.Service, app *fiber.App) handler {

	handler := handler{
		service: service,
	}

	app.Get("/", handler.HelloWorld)
	app.Get("/expense", handler.GetAllExpenses)
	app.Post("/expense", handler.AddExpense)

	return handler
}

func (h handler) HelloWorld(c fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", c.Params("*"))
	err := c.SendString(msg) // => ✋ register

	if err != nil {
		panic("")
	}

	return nil
}

func (h handler) GetAllExpenses(c fiber.Ctx) error {
	// location, _ := time.LoadLocation("America/Sao_Paulo")

	// model := model.Expense{
	// 	ID:           "2",
	// 	Value:        10.00,
	// 	PaymentDate:  time.Now().In(location),
	// 	Installments: 1,
	// }

	// err := h.service.InsertOne(c.Context(), model)

	// if err != nil {
	// 	panic("")
	// }

	return nil
}

func (h handler) AddExpense(c fiber.Ctx) error {
	var body model.Expense

	json.Unmarshal(c.Body(), &body)

	h.service.InsertOne(c.Context(), body)

	return nil
}
