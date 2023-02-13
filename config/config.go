package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config .
	Config struct {
		HTTP         `yaml:"http"`
		NatsStreamig `yaml:"nats_straming"`
		PG
	}

	// HTTP .
	HTTP struct {
		Port string `yaml:"port"`
	}

	// PG .
	PG struct {
		URL string
	}

	// NatsStreamig .
	NatsStreamig struct {
		ClusterID string `yaml:"cluster_id"`
		ClientID  string `yaml:"client_id"`
		Subject   string `yaml:"subject"`
		URL       string `env:"NATS_URL"`
	}
)

// NewConfig .
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	cfg.PG.URL = databaseURL()

	return cfg, nil
}

// DatabaseURL ...
func databaseURL() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
}
