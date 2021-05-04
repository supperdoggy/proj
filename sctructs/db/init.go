package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*DB, error) {
	creds := CredsForDB{
		Name:     "users",
		Host:     "localhost",
		Port:     "5432",
		User:     "maks",
		Password: "abc123",
	}

	db, err := gorm.Open(postgres.Open(creds.GetURI()), &gorm.Config{})
	if err != nil || db == nil {
		return nil, err
	}
	return &DB{D: db}, nil
}
