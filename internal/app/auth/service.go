package auth

import (
	"context"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/StratuStore/auth/internal/app/tokens"
	"github.com/mbretter/go-mongodb/types"
	"log/slog"
)

type Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Storage interface {
	GetUser(ctx context.Context, sub string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) error
	AddSession(ctx context.Context, session *core.Session) error
	UpdateUser(ctx context.Context, user *core.User) error
	UpdateSession(ctx context.Context, session *core.Session) error
	DeleteSession(ctx context.Context, sessionID types.ObjectId) error
	GetSession(ctx context.Context, id types.ObjectId) (*core.Session, error)
}

type Context interface {
	context.Context
	UserAgent() string
}

type Service struct {
	l                *slog.Logger
	storage          Storage
	accessTokenAuth  *tokens.AccessJWTWorker
	refreshTokenAuth *tokens.RefreshJWTWorker
}

func New(logger *slog.Logger, storage Storage, accessTokenAuth *tokens.AccessJWTWorker, refreshTokenAuth *tokens.RefreshJWTWorker) *Service {
	return &Service{
		l:                logger.With(slog.String("module", "internal.app.auth")),
		storage:          storage,
		accessTokenAuth:  accessTokenAuth,
		refreshTokenAuth: refreshTokenAuth,
	}
}
