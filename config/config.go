package config

import (
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
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
