package config

import (
	"fmt"
	"os"

	"github.com/fc1g/http-server/internal/validation"
)

type Secrets struct {
	Environment string `validate:"required,oneof=development test staging production"`
}

func LoadSecrets(v validation.Validator) (*Secrets, error) {
	secrets := &Secrets{
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	if err := v.Struct(secrets); err != nil {
		return nil, fmt.Errorf("invalid secrets: %w", err)
	}

	return secrets, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
