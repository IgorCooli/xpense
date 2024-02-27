package user

import (
	"context"
	"encoding/json"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/business/service/user"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service user.Service
}

func NewHandler(ctx context.Context, service user.Service, app *fiber.App) handler {

	handler := handler{
		service: service,
	}

	app.Post("/user/register", handler.RegisterUser)

	return handler
}

func (h handler) RegisterUser(c fiber.Ctx) error {
	var body model.User
	json.Unmarshal(c.Body(), &body)

	h.service.RegisterUser(c.Context(), body)

	return nil
}
