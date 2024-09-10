package config

import (
	"os"
)

type Config struct {
    DatabaseURL  string
    RabbitMQURL  string
}

func LoadConfig() *Config {
    return &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        RabbitMQURL: os.Getenv("RABBITMQ_URL"),
    }
}