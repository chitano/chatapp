package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Port     string
	Host     string
	Database string
	User     string
	Password string
}

func LoadConfig() *DatabaseConfig {
	//Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &DatabaseConfig{
		Port:     os.Getenv("PORT"),
		Host:     os.Getenv("HOST"),
		Database: os.Getenv("DATABASE"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
	}
}

func ConnectDatabase(cfg *DatabaseConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
}
