package auth

import (
	"context"
	"github.com/StratuStore/auth/internal/libs/errors"
	"log/slog"
)

func (s *Service) Revoke(ctx context.Context, requestBody RefreshRequest) error {
	l := s.l.With(
		slog.String("op", "internal.auth.Revoke"),
	)

	_, session, err := s.validateRefreshToken(ctx, requestBody.RefreshToken)
	if err != nil {
		return err
	}

	err = s.storage.DeleteSession(ctx, session.ID)
	if err != nil {
		return errors.NewInternalError(l, "failed to delete session", err)
	}

	return nil
}
