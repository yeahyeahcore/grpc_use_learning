package utils

import (
	"fmt"

	envParser "github.com/caarlos0/env/v6"
	envLoader "github.com/joho/godotenv"
)

// ParseENV parsing default env file and returns struct pointer by generic
func ParseENV[T interface{}]() (*T, error) {
	var configuration T

	if err := envLoader.Load(".env"); err != nil {
		return nil, fmt.Errorf("loading env error %s", err.Error())
	}
	if err := envParser.Parse(&configuration); err != nil {
		return nil, fmt.Errorf("parse env error %s", err.Error())
	}

	return &configuration, nil
}
