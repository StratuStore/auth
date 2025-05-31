package google

import (
	"github.com/StratuStore/auth/internal/app/auth"
	"github.com/StratuStore/auth/internal/libs/config"
	"log/slog"
)

type Service struct {
	auth           *auth.Service
	l              *slog.Logger
	googleClientID string
}

func New(logger *slog.Logger, auth *auth.Service, cfg *config.Config) *Service {
	return &Service{
		l:              logger.With(slog.String("module", "internal.app.auth.google")),
		auth:           auth,
		googleClientID: cfg.GoogleClientID,
	}
}
