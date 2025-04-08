package config

import (
	"os"

	"github.com/fogtz/goVacations.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if db already exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create the database file
		logger.Info("Database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Creating DB and Connecting
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to SQLite database: %v", err)
		return nil, err
	}
	// Migrating the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating SQLite database: %v", err)
		return nil, err
	}

	// Return DB if everything is ok
	return db, nil
}
