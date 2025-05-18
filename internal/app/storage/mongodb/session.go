package mongodb

import (
	"context"
	"fmt"
	"github.com/StratuStore/auth/internal/app/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) AddSession(ctx context.Context, session *core.Session) error {
	db := s.db

	result, err := db.Collection("sessions").
		InsertOne(ctx, bson.D{{"userSub", session.UserSub}, {"salt", session.Salt}, {"deviceData", session.DeviceData}})
	if err != nil {
		return err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil
	}

	session.ID = id.Hex()

	return err
}

func (s *Storage) UpdateSession(ctx context.Context, session *core.Session) error {
	db := s.db

	objID, err := primitive.ObjectIDFromHex(session.ID)
	if err != nil {
		return fmt.Errorf("unable to convert sessionID to objectID: %w", err)
	}

	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", bson.D{{"salt", session.Salt}, {"deviceData", session.DeviceData}}}}
	_, err = db.Collection("sessions").
		UpdateOne(
			ctx,
			filter,
			update,
		)

	return err
}

func (s *Storage) DeleteSession(ctx context.Context, id string) error {
	db := s.db

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("unable to convert sessionID to objectID: %w", err)
	}

	filter := bson.D{{"_id", objID}}
	result := db.Collection("sessions").
		FindOneAndDelete(
			ctx,
			filter,
		)

	return result.Err()
}

func (s *Storage) GetSession(ctx context.Context, id string) (*core.Session, error) {
	db := s.db

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("unable to convert sessionID to objectID: %w", err)
	}

	var session core.Session
	err = db.Collection("sessions").
		FindOne(ctx, bson.M{"_id": objID}).
		Decode(&session)

	return &session, err
}
