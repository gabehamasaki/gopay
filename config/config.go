package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error initializing dotenv: %v", err.Error())
	}

	// Initialize database
	db, err = InicializeDatabase()
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err.Error())
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
