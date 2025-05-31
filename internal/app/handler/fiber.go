package handler

import (
	"context"
	"github.com/StratuStore/auth/internal/app/auth"
	"github.com/StratuStore/auth/internal/app/auth/firebase"
	"github.com/StratuStore/auth/internal/app/auth/google"
	"github.com/StratuStore/auth/internal/libs/config"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log/slog"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Handler struct {
	app             *fiber.App
	l               *slog.Logger
	cfg             *config.Config
	googleService   *google.Service
	firebaseService *firebase.Service
	service         *auth.Service
}

func New(l *slog.Logger, cfg *config.Config, googleService *google.Service, firebaseService *firebase.Service, service *auth.Service) *Handler {
	h := &Handler{
		l:   l.With(slog.String("module", "internal.auth.handler")),
		cfg: cfg,
		app: fiber.New(fiber.Config{
			IdleTimeout:  cfg.IdleTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		}),
		googleService:   googleService,
		firebaseService: firebaseService,
		service:         service,
	}

	h.Register()

	return h
}

func (h *Handler) Register() {
	h.registerDefaults()

	h.app.Post("/google", h.Google)
	h.app.Post("/firebase", h.Firebase)
	h.app.Post("/refresh", h.Refresh)
	h.app.Delete("/revoke", h.Revoke)
}

func (h *Handler) registerDefaults() {
	if h.cfg.Env == "dev" {
		h.app.Use(cors.New(cors.ConfigDefault))
	} else {
		h.app.Use(cors.New(cors.Config{
			AllowOrigins: h.cfg.CORSOrigins,
		}))
	}

	h.app.Use(recover.New(recover.Config{
		StackTraceHandler: func(c *fiber.Ctx, r any) {
			h.l.Error("fiber panicked", slog.Any("err", r))
		},
	}))

	h.app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true // TODO: service check
		},
		ReadinessEndpoint: "/ready",
	}))
}

func (h *Handler) Start(_ context.Context) error {
	l := h.l.With("op", "internal.config-manager.handler.Start")

	addr := net.JoinHostPort(h.cfg.Host, h.cfg.Port)

	go func() {
		if err := h.app.Listen(addr); err != nil {
			l.Error("server error", slog.String("err", err.Error()))
		}
	}()

	return nil
}

func (h *Handler) Stop(ctx context.Context) error {
	return h.app.ShutdownWithContext(ctx)
}
