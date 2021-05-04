package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/supperdoggy/score/sctructs"
)

func Init() (*sctructs.DB, error) {
	creds := sctructs.CredsForDB{
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
	return &sctructs.DB{D: db}, nil
}
