package firebase

import (
	"context"
	"github.com/StratuStore/auth/internal/app/auth"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/StratuStore/auth/internal/libs/errors"
	"log/slog"
)

type LoginData struct {
	JWT string `json:"firebaseJWT"`
}

type Context interface {
	context.Context
	UserAgent() string
}

func (s *Service) Authenticate(ctx Context, loginData LoginData) (auth.Response, error) {
	l := s.l.With(
		slog.String("op", "CreateSession"),
	)

	client, err := s.fb.Auth(ctx)
	if err != nil {
		return auth.Response{}, errors.NewValidationError(
			l,
			"unable to get firebase client",
			"service temporary unavailable",
			err,
		)
	}

	l.Debug("processing jwt from CreateSession")
	payload, err := client.VerifyIDTokenAndCheckRevoked(ctx, loginData.JWT)
	if err != nil {
		return auth.Response{}, errors.NewValidationError(
			l,
			"unable to parse jwt from CreateSession",
			"invalid access token",
			err,
		)
	}

	l.Debug("processing claims")
	claimsUser, err := core.NewUserFromClaims(payload.Claims)
	if err != nil {
		return auth.Response{}, errors.NewInternalError(
			l,
			"unable to parse claims from CreateSession",
			err,
		)
	}

	return s.auth.CreateSession(ctx, claimsUser)
}
