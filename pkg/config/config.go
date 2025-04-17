package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Schema struct {
	DatabaseURI string `env:"CONNECTION_STRING" mapstructure:"connection_string"`
	Environment string `env:"ENVIRONMENT" mapstructure:"environment"`
	HttpPort    string `env:"HTTP_PORT" mapstructure:"http_port"`
}

var (
	cfg Schema
)

func LoadConfig() *Schema {
	environment := os.Getenv("ENVIRONMENT")
	switch environment {
	case "PRODUCTION":
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Error loading.env file: %v", err)
		}

	default:
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatalf("Error loading.env file: %v", err)
		}
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing env file: %v", err)
	}

	return &cfg
}

func GetConfig() *Schema {
	return &cfg
}
