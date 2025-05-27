package mongodb

import (
	"context"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/mbretter/go-mongodb/types"
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

	session.ID = types.ObjectId(id.Hex())

	return err
}

func (s *Storage) UpdateSession(ctx context.Context, session *core.Session) error {
	db := s.db

	filter := bson.D{{"_id", session.ID}}
	update := bson.D{{"$set", bson.D{{"salt", session.Salt}, {"deviceData", session.DeviceData}}}}
	_, err := db.Collection("sessions").
		UpdateOne(
			ctx,
			filter,
			update,
		)

	return err
}

func (s *Storage) DeleteSession(ctx context.Context, id types.ObjectId) error {
	db := s.db

	filter := bson.D{{"_id", id}}
	result := db.Collection("sessions").
		FindOneAndDelete(
			ctx,
			filter,
		)

	return result.Err()
}

func (s *Storage) GetSession(ctx context.Context, id types.ObjectId) (*core.Session, error) {
	db := s.db

	var session core.Session
	err := db.Collection("sessions").
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&session)

	return &session, err
}
