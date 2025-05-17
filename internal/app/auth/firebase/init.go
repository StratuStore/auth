package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/StratuStore/auth/internal/app/auth"
	"github.com/StratuStore/auth/internal/libs/config"
	"google.golang.org/api/option"
	"log/slog"
)

type Service struct {
	auth *auth.Service
	l    *slog.Logger
	fb   *firebase.App
}

func New(logger *slog.Logger, auth *auth.Service, cfg *config.Config) (*Service, error) {
	opt := option.WithCredentialsJSON([]byte(cfg.FirebaseConfig))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return &Service{
		l:    logger.With(slog.String("module", "internal.app.auth.firebase")),
		auth: auth,
		fb:   app,
	}, nil
}
