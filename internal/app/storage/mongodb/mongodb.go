package mongodb

import (
	"context"
	"github.com/StratuStore/auth/internal/libs/config"
	"github.com/cenkalti/backoff/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	db *mongo.Database
}

func New(cfg *config.Config) *Storage {
	client, err := openConnection(cfg.MongoConnectionString(), cfg.MongoMaxRetries)
	if err != nil {
		panic(err)
	}

	return &Storage{client.Database(cfg.MongoDB)}

}
func openConnection(connectionString string, maxRetries uint) (*mongo.Client, error) {
	operation := func() (*mongo.Client, error) {
		return mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	}

	return backoff.Retry(
		context.Background(),
		operation,
		backoff.WithBackOff(backoff.NewExponentialBackOff()),
		backoff.WithMaxTries(maxRetries),
	)
}
