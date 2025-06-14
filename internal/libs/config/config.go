package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Firebase struct {
	FirebaseConfig string `env:"FIREBASE_CONFIG"`
}

type Google struct {
	GoogleClientID          string `env:"GOOGLE_CLIENT_ID"`
	AccessTokenExpInMinutes int    `env:"ACCESS_TOKEN_EXP_IN_MINUTES" env-default:"15"`
	RefreshTokenExpInDays   int    `env:"REFRESH_TOKEN_EXP_IN_DAYS" env-default:"7"`
	AuthSecret              string `env:"AUTH_SECRET"`
}

type Mongo struct {
	MongoUser       string `env:"MONGO_USER" env-default:"root"`
	MongoPass       string `env:"MONGO_PASS" env-default:"password"`
	MongoHost       string `env:"MONGO_HOST" env-default:"localhost"`
	MongoPort       string `env:"MONGO_PORT" env-default:"27017"`
	MongoDB         string `env:"MONGO_DB" env-default:"auth"`
	MongoMaxRetries uint   `env:"MONGO_MAX_RETRIES" env-default:"5"`
}

func (m *Mongo) MongoConnectionString() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v", m.MongoUser, m.MongoPass, m.MongoHost, m.MongoPort)
}

type Handler struct {
	Host         string        `env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port         string        `env:"HTTP_PORT" env-default:"8080"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" env-default:"10s"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" env-default:"10s"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT" env-default:"10s"`
	CORSOrigins  string        `env:"HTTP_CORS_ORIGINS"`
}

type Logger struct {
	Level string `env:"LOGGER_LEVEL" env-default:"INFO"`
}

type Config struct {
	Firebase
	Google
	Mongo
	Logger
	Handler
	Env string `env:"ENV" env-default:"dev"`
}

const filepath = "./.env"

func New() (*Config, error) {
	var c Config

	err := cleanenv.ReadConfig(filepath, &c)
	if errors.Is(err, os.ErrNotExist) {
		err = cleanenv.ReadEnv(&c)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to read config: %w", err)
	}

	return &c, nil
}
