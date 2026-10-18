package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	return db, nil

}
