package api

import (
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

// Config defines application configuration.
//
// Environment variables:
//   - APP_PORT     (default: "8000")
//   - SERVICE_NAME (default: "Bookmark-service")
//   - INSTANCE_ID  (default: auto-generated UUID)
type Config struct {
	AppPort     string `default:"8000" envconfig:"APP_PORT"`
	ServiceName string `default:"Bookmark-service" envconfig:"SERVICE_NAME"`
	InstanceID  string ` envconfig:"INSTANCE_ID"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("api", cfg)

	if err != nil {
		return nil, err
	}

	if cfg.InstanceID == "" {
		cfg.InstanceID = uuid.New().String()
	}

	return cfg, err
}
