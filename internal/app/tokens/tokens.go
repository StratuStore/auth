package tokens

import (
	"github.com/StratuStore/auth/internal/libs/config"
	"time"
)

type AccessJWTWorker struct {
	*JWTWorker
}

func NewAccessJWTWorker(cfg *config.Config) *AccessJWTWorker {
	return &AccessJWTWorker{
		NewJWTWorker(cfg.AuthSecret, time.Duration(cfg.AccessTokenExpInMinutes)*time.Minute),
	}
}

type RefreshJWTWorker struct {
	*JWTWorker
}

func NewRefreshJWTWorker(cfg *config.Config) *RefreshJWTWorker {
	return &RefreshJWTWorker{
		NewJWTWorker(cfg.AuthSecret, time.Duration(cfg.RefreshTokenExpInDays)*time.Hour*24),
	}
}
