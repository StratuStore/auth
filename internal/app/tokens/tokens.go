package tokens

import (
	"github.com/StratuStore/auth/internal/libs/config"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"time"
)

type AccessJWTWorker struct {
	*jwtauth.JWTAuth
}

func NewAccessJWTWorker(cfg *config.Config) *AccessJWTWorker {
	return &AccessJWTWorker{
		jwtauth.New(
			"HS512",
			[]byte(cfg.AuthSecret),
			nil,
			jwt.WithAcceptableSkew(time.Duration(cfg.AccessTokenExpInMinutes)*time.Minute),
		),
	}
}

type RefreshJWTWorker struct {
	*jwtauth.JWTAuth
}

func NewRefreshJWTWorker(cfg *config.Config) *RefreshJWTWorker {
	return &RefreshJWTWorker{
		jwtauth.New(
			"HS512",
			[]byte(cfg.AuthSecret),
			nil,
			jwt.WithAcceptableSkew(time.Duration(cfg.RefreshTokenExpInDays)*time.Hour*24),
		),
	}
}
