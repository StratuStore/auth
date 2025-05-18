package auth

import (
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/StratuStore/auth/internal/libs/errors"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) CreateSession(ctx Context, claimsUser core.User) (Response, error) {
	l := s.l.With(
		slog.String("op", "CreateSession"),
	)

	l.Debug("checking if user already exists")
	user, err := s.storage.GetUser(ctx, claimsUser.Sub)
	if err != nil {
		if err = s.storage.AddUser(ctx, &claimsUser); err != nil {
			return Response{}, errors.NewInternalError(l, "unable to add user to db", err)
		}
		user = &claimsUser
	} else {
		l.Debug("updating user data")
		user.Name = claimsUser.Name
		user.Picture = claimsUser.Picture

		err = s.storage.UpdateUser(ctx, user)
		if err != nil {
			return Response{}, errors.NewInternalError(l, "unable to update user data", err)
		}
	}

	l.Debug("creating session")
	session := &core.Session{
		UserSub:    user.Sub,
		Salt:       uuid.New(),
		DeviceData: ctx.UserAgent(),
	}
	err = s.storage.AddSession(ctx, session)
	if err != nil {
		return Response{}, errors.NewInternalError(l, "unable to add session to db", err)
	}

	return s.generateTokens(user, session)
}

func (s *Service) generateTokens(user *core.User, session *core.Session) (Response, error) {
	l := s.l.With(
		slog.String("op", "generateTokens"),
	)

	l.Debug("generating refresh token")
	refreshTokenString, err := s.refreshTokenAuth.Encode(session.GetClaims())
	if err != nil {
		return Response{}, errors.NewInternalError(l, "unable to encode refresh token", err)
	}

	l.Debug("generating access token")
	accessTokenString, err := s.accessTokenAuth.Encode(user.GetClaims())
	if err != nil {
		return Response{}, errors.NewInternalError(l, "unable to generate access token", err)
	}

	return Response{AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}
