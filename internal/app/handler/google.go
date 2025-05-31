package handler

import (
	"context"
	"github.com/StratuStore/auth/internal/app/auth/google"
	"github.com/StratuStore/auth/internal/app/core"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Google(c *fiber.Ctx) error {
	l := h.l.With(slog.String("op", "Google"))

	c.Accepts("application/json")

	var data google.LoginData
	if err := c.BodyParser(&data); err != nil {
		l.Debug("unable to parse request body", slog.String("err", err.Error()))

		return c.Status(http.StatusBadRequest).JSON(core.NewErrorResponse("unable to parse request body"))
	}

	entity, err := h.googleService.Authenticate(&CustomContext{
		Context:   c.Context(),
		userAgent: c.Get("User-Agent"),
	}, data)
	if err != nil {
		return ProcessError(l, c, err)
	}

	l.Debug("google authenticated")

	return c.Status(http.StatusOK).JSON(core.NewOKResponse(entity))
}

type CustomContext struct {
	context.Context
	userAgent string
}

func (c *CustomContext) UserAgent() string {
	return c.userAgent
}
