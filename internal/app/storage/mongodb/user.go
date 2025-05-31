package mongodb

import (
	"context"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/mbretter/go-mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) GetUser(ctx context.Context, sub string) (*core.User, error) {
	db := s.db

	var user core.User
	err := db.Collection("users").
		FindOne(ctx, bson.M{"sub": sub}).
		Decode(&user)

	return &user, err
}

func (s *Storage) AddUser(ctx context.Context, user *core.User) error {
	db := s.db

	result, err := db.Collection("users").
		InsertOne(ctx, bson.D{{"sub", user.Sub}, {"email", user.Email}, {"name", user.Name}, {"picturePath", user.Picture}})
	if err != nil {
		return err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil
	}

	user.ID = types.ObjectId(id.Hex())

	return err
}

func (s *Storage) UpdateUser(ctx context.Context, user *core.User) error {
	db := s.db

	filter := bson.D{{"sub", user.Sub}}
	update := bson.D{{"$set", bson.D{{"email", user.Email}, {"name", user.Name}, {"picturePath", user.Picture}}}}
	_, err := db.Collection("users").
		UpdateOne(
			ctx,
			filter,
			update,
		)

	return err
}
