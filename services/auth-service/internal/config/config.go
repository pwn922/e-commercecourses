package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL  string
    //RabbitMQURL  string
}

func LoadConfig() *Config {
    // Carga las variables del archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    return &Config{
        DatabaseURL: os.Getenv("DATABASE_AUTH_URL"),
//        RabbitMQURL: os.Getenv("RABBITMQ_URL"),
    }
}