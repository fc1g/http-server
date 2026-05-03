package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Environment string

const (
	Development Environment = "development"
	Test        Environment = "test"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

type ServerConfig struct {
	Network string
	Addr    int
}

type Config struct {
	Environment
	Server ServerConfig
}

func parseEnvironment(s string) (Environment, error) {
	switch env := Environment(s); env {
	case Development, Test, Staging, Production:
		return env, nil
	default:
		return "", fmt.Errorf("unknown environment %q", s)
	}
}

func Load() (*Config, error) {
	env, err := parseEnvironment(getEnv("ENVIRONMENT", string(Development)))
	if err != nil {
		return nil, err
	}

	return &Config{
		Environment: env,
		Server: ServerConfig{
			Network: getEnv("NETWORK", "tcp"),
			Addr:    getEnvInt("ADDR", 3000),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Failed to convert %s: %v", key, err)
		return fallback
	}

	return val
}
