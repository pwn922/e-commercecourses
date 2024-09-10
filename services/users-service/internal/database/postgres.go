package database

import (
	"log"

	"github.com/pwn922/users-service/internal/config"
	"github.com/pwn922/users-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var database *gorm.DB

func InitDatabase(cfg *config.Config) {
    dns := cfg.DatabaseURL
    if dns == "" {
        log.Fatal("Database URL is not provided in the configuration.")
    }

    log.Printf("Database URL: %s", dns)

    var err error
    database, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

func GetDB() *gorm.DB {
    if database == nil {
        log.Fatal("Database has not been initialized. Call InitDatabase first.")
    }
    return database
}

func Migrate() {
    if database == nil {
        log.Fatal("Database has not been initialized. Call InitDatabase first.")
    }

    if err := database.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("Failed to automigrate database: %v", err)
    }
}

func CloseDatabase() {
    if database == nil {
        log.Fatal("Database has not been initialized. Call InitDatabase first.")
    }

    sqlDB, err := database.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }
    if err := sqlDB.Close(); err != nil {
        log.Fatalf("Failed to close database connection: %v", err)
    } 
}
