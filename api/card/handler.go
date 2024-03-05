package card

import (
	"context"
	"encoding/json"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/business/service/card"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service card.Service
}

func NewHandler(ctx context.Context, service card.Service, app *fiber.App) handler {

	handler := handler{
		service: service,
	}

	app.Get("/card/:id", handler.FindById)
	app.Post("/card/register", handler.RegisterCard)

	return handler
}

func (h handler) FindById(c fiber.Ctx) error {
	cardId := c.Params("id")

	response, err := h.service.FindById(c.Context(), cardId)

	if err != nil {
		panic("Could not find card")
	}

	if response.UserID != extractUserId(c.GetReqHeaders()) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Operation forbidden"})
	}

	c.JSON(response)

	return nil
}

func (h handler) RegisterCard(c fiber.Ctx) error {
	var body model.Card
	json.Unmarshal(c.Body(), &body)

	if body.UserID != extractUserId(c.GetReqHeaders()) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Operation forbidden"})
	}

	h.service.RegisterCard(c.Context(), body)

	return nil
}

func extractUserId(headers map[string][]string) string {
	header := headers["X-User-Id"]
	headerValue := header[0]

	return headerValue
}
