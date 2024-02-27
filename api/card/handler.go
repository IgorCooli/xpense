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

	app.Post("/card/register", handler.RegisterCard)

	return handler
}

func (h handler) RegisterCard(c fiber.Ctx) error {
	var body model.Card
	json.Unmarshal(c.Body(), &body)

	h.service.RegisterCard(c.Context(), body)

	return nil
}
