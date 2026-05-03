package config

import (
	"fmt"

	"github.com/fc1g/http-server/internal/validation"
	"github.com/spf13/viper"
)

type Connection struct {
	ReadTimeout  int `mapstructure:"read_timeout" validate:"required,min=15"`
	WriteTimeout int `mapstructure:"write_timeout" validate:"required,min=15"`
	BufferSize   int `mapstructure:"buffer_size" validate:"required,min=4096"`
}

type Server struct {
	Network    string `mapstructure:"network" validate:"required"`
	Addr       int    `mapstructure:"addr" validate:"required"`
	Connection `mapstructure:"connection" validate:"required"`
}

type Settings struct {
	Server `mapstructure:"server" validate:"required"`
}

func Load(v validation.Validator) (*Settings, error) {
	settings := &Settings{}
	if err := viper.Unmarshal(settings); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if err := v.Struct(settings); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return settings, nil
}
