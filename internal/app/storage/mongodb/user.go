package mongodb

import (
	"context"
	"github.com/StratuStore/auth/internal/app/core"
	"github.com/mbretter/go-mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*core.User, error) {
	db := s.db

	var user core.User
	err := db.Collection("users").
		FindOne(ctx, bson.M{"email": email}).
		Decode(&user)

	return &user, err
}

func (s *Storage) GetUserByID(ctx context.Context, id string) (*core.User, error) {
	db := s.db

	idHex, err := types.ObjectIdFromHex(id)
	if err != nil {
		return nil, err
	}

	var user core.User
	err = db.Collection("users").
		FindOne(ctx, bson.M{"_id": idHex}).
		Decode(&user)

	return &user, err
}

func (s *Storage) AddUser(ctx context.Context, user *core.User) error {
	db := s.db

	result, err := db.Collection("users").
		InsertOne(ctx, bson.D{{"email", user.Email}, {"name", user.Name}, {"picturePath", user.Picture}})
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

	filter := bson.D{{"email", user.Email}}
	update := bson.D{{"$set", bson.D{{"email", user.Email}, {"name", user.Name}, {"picturePath", user.Picture}}}}
	_, err := db.Collection("users").
		UpdateOne(
			ctx,
			filter,
			update,
		)

	return err
}
