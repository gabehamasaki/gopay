package config

import (
	"os"

	"github.com/gabehamasaki/gopay/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InicializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("database")

	DSN := os.Getenv("DATABASE_DSN")
	if DSN == "" {
		DSN = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	}

	// Connection to database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		logger.Errorf("database opening error: %v", err.Error())
		return nil, err
	}

	// Migration schemas
	err = db.AutoMigrate(&models.Account{})
	if err != nil {
		logger.Errorf("database automigration error: %v", err.Error())
		return nil, err
	}

	// Return database
	return db, nil
}
