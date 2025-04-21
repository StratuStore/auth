package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

func NewConfig[T any](filepath string, cfg *T) error {
	return cleanenv.ReadConfig(filepath, &cfg)
}
