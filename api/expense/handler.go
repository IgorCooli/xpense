package expense

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/business/service/expense"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service expense.Service
}

func NewHandler(ctx context.Context, service expense.Service, app *fiber.App) handler {

	handler := handler{
		service: service,
	}

	app.Get("/", handler.HelloWorld)
	app.Get("/expense/search", handler.SearchExpenses)
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

func (h handler) SearchExpenses(c fiber.Ctx) error {

	month := c.Query("month")
	year := c.Query("year")

	if month == "" || year == "" {
		panic("The query params are not complete")
	}

	result := h.service.Search(c.Context(), month, year)

	c.JSON(result)
	return nil
}

func (h handler) AddExpense(c fiber.Ctx) error {
	var body model.Expense
	json.Unmarshal(c.Body(), &body)

	h.service.AddExpense(c.Context(), body)

	return nil
}
