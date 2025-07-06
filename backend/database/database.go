package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

var DB *gorm.DB

func DefaultConfig() *Config {
	return &Config{
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnv("POSTGRES_DB", "smashfriend"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func Connect(config *Config) (*gorm.DB, error) {
	dsn := "host=" + config.Host +
		" user=" + config.User +
		" password=" + config.Password +
		" dbname=" + config.DBName +
		" port=" + config.Port +
		" sslmode=" + config.SSLMode

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = database

	log.Println("Successfully connected to database")
	return DB, nil
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
