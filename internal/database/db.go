package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("error migrating database: %v", err)
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate()
}
