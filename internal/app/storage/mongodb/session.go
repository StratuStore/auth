package mongodb

import (
	"context"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) AddSession(ctx context.Context, session *core.Session) error {
	db := s.db

	_, err := db.Collection("sessions").
		InsertOne(ctx, session)

	return err
}

func (s *Storage) UpdateSession(ctx context.Context, session *core.Session) error {
	db := s.db

	filter := bson.D{{"_id", session.ID}}
	update := bson.D{{"$set", bson.D{{"refreshToken", session.RefreshToken}, {"deviceData", session.DeviceData}}}}
	_, err := db.Collection("sessions").
		UpdateOne(
			ctx,
			filter,
			update,
		)

	return err
}

func (s *Storage) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	db := s.db

	filter := bson.D{{"_id", sessionID}}
	result := db.Collection("sessions").
		FindOneAndDelete(
			ctx,
			filter,
		)

	return result.Err()
}
