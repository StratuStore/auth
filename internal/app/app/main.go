package app

import (
	"github.com/StratuStore/auth/internal/app/auth"
	"github.com/StratuStore/auth/internal/app/auth/firebase"
	"github.com/StratuStore/auth/internal/app/auth/google"
	"github.com/StratuStore/auth/internal/app/handler"
	"github.com/StratuStore/auth/internal/app/storage/mongodb"
	"github.com/StratuStore/auth/internal/app/tokens"
	"github.com/StratuStore/auth/internal/libs/config"
	"github.com/StratuStore/auth/internal/libs/log"
	"go.uber.org/fx"
)

func CreateApp(cfg *config.Config) fx.Option {
	return fx.Options(
		fx.Supply(
			cfg,
		),
		fx.Provide(
			log.New,

			fx.Annotate(mongodb.New, fx.As(new(auth.Storage))),
			tokens.NewAccessJWTWorker,
			tokens.NewRefreshJWTWorker,
			auth.New,
			google.New,
			firebase.New,

			handler.New,
		),
		fx.Invoke(
			startHTTPServer,
		),
	)
}

func startHTTPServer(lifecycle fx.Lifecycle, h *handler.Handler) {
	lifecycle.Append(fx.Hook{
		OnStart: h.Start,
		OnStop:  h.Stop,
	})
}
