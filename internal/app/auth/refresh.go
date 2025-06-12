package auth

import (
	"context"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/StratuStore/auth/internal/libs/errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"log/slog"
)

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (s *Service) Refresh(ctx Context, requestBody RefreshRequest) (Response, error) {
	l := s.l.With(
		slog.String("op", "internal.auth.Refresh"),
	)

	user, session, err := s.validateRefreshToken(ctx, requestBody.RefreshToken)
	if err != nil {
		return Response{}, err
	}

	session.Salt = uuid.New()
	session.DeviceData = ctx.UserAgent()
	err = s.storage.UpdateSession(ctx, session)
	if err != nil {
		return Response{}, errors.NewInternalError(l, "failed to update session")
	}

	return s.generateTokens(user, session)
}

func (s *Service) validateRefreshToken(ctx context.Context, refreshToken string) (*core.User, *core.Session, error) {
	l := s.l.With(
		slog.String("op", "internal.auth.validateRefreshToken"),
	)

	token, err := jwtauth.VerifyToken(s.refreshTokenAuth.JWTAuth, refreshToken)
	if err != nil {
		return nil, nil, errors.NewValidationError(l, "failed to decode refresh token", "invalid refresh token", err)
	}

	sessionClaims := token.PrivateClaims()
	sessionClaims["jti"] = token.JwtID()
	sessionClaims["sub"] = token.Subject()
	claimsSession, err := core.NewSessionFromClaims(sessionClaims)
	if err != nil {
		return nil, nil, errors.NewValidationError(l, "failed to decode session", "invalid refresh token", err)
	}

	user, err := s.storage.GetUserByID(ctx, claimsSession.UserSub)
	if err != nil {
		return nil, nil, errors.NewNotFoundError(l, "failed to get user", "user is not found", err)
	}

	session, err := s.storage.GetSession(ctx, claimsSession.ID)
	if err != nil || session.UserSub != string(user.ID) {
		return nil, nil, errors.NewNotFoundError(l, "session is not found", "session is not found", err)
	}
	if session.Salt != claimsSession.Salt {
		return nil, nil, errors.NewValidationError(l, "invalid refresh token", "invalid refresh token")
	}

	return user, session, nil
}
