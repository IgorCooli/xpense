package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/IgorCooli/xpense/internal/business/model/request"
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

	app.Post("/auth", handler.Authenticate)

	return handler
}

func (h handler) Authenticate(c fiber.Ctx) error {
	var body request.Credentials
	json.Unmarshal(c.Body(), &body)

	token, err := h.service.AuthenticateUser(c.Context(), body)

	if err != nil {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}

	type response struct {
		Token string `json:"token"`
	}

	c.JSON(response{Token: token})
	return nil
}
