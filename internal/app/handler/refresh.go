package handler

import (
	"github.com/StratuStore/auth/internal/app/auth"
	"github.com/StratuStore/auth/internal/app/core"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Refresh(c *fiber.Ctx) error {
	l := h.l.With(slog.String("op", "Refresh"))

	c.Accepts("application/json")

	var data auth.RefreshRequest
	if err := c.BodyParser(&data); err != nil {
		l.Debug("unable to parse request body", slog.String("err", err.Error()))

		return c.Status(http.StatusBadRequest).JSON(core.NewErrorResponse("unable to parse request body"))
	}

	entity, err := h.service.Refresh(&CustomContext{
		Context:   c.Context(),
		userAgent: c.Get("User-Agent"),
	}, data)
	if err != nil {
		return ProcessError(l, c, err)
	}

	l.Debug("refreshed")

	return c.Status(http.StatusOK).JSON(core.NewOKResponse(entity))
}
