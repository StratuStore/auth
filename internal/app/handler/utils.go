package handler

import (
	"errors"
	"github.com/StratuStore/auth/internal/app/core"
	"log/slog"
	"net/http"
	"strconv"

	ownerrors "github.com/StratuStore/auth/internal/libs/errors"
	"github.com/gofiber/fiber/v2"
)

func ProcessError(l *slog.Logger, c *fiber.Ctx, err error) error {
	var userErr ownerrors.UserError
	if errors.As(err, &userErr) {
		l.Debug("unable to execute service", slog.String("err", err.Error()))

		return c.Status(userErr.Status()).JSON(core.NewResponseByUserError(userErr))
	}
	l.Error("unable to execute service (unknown error)", slog.String("err", err.Error()))

	return c.Status(http.StatusInternalServerError).JSON(core.NewErrorResponse("internal error"))
}

func Query(l *slog.Logger, c *fiber.Ctx, key string, errs ...error) (int, error) {
	dataStr := c.Query(key, "0")
	data, err := strconv.Atoi(dataStr)
	if err != nil {
		l.Debug("unable to parse data", slog.String("key", key), slog.String("err", err.Error()))

		return 0, ownerrors.New(errors.Join(errs...), "unable to parse data", "unable to parse "+key, http.StatusBadRequest)
	}

	return data, nil
}
