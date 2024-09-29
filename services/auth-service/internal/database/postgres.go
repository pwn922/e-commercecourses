package database

import (
	"log"

	"github.com/pwn922/auth-service/internal/config"
	"github.com/pwn922/auth-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgDb *gorm.DB

func InitDatabase(cfg *config.Config) {
	dns := cfg.DatabaseURL
	if dns == "" {
		log.Fatal("Database URL is not provided in the configuration.")
	}

	log.Printf("Database URL: %s", dns)

	var err error
	pgDb, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func GetDB() *gorm.DB {
	if pgDb == nil {
		log.Fatal("Database has not been initialized. Call InitDatabase first.")
	}
	return pgDb
}

func Migrate() {
	if pgDb == nil {
		log.Fatal("Database has not been initialized. Call InitDatabase first.")
	}

	if err := pgDb.AutoMigrate(&models.User{}, &models.Role{}); err != nil {
		log.Fatalf("Failed to automigrate pgDb: %v", err)
	}
}

func CloseDatabase() {
	if pgDb == nil {
		log.Fatal("Database has not been initialized. Call InitDatabase first.")
	}

	sqlDB, err := pgDb.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}