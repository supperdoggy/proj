package db

import (
	"github.com/supperdoggy/score/sctructs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitUsersDB() (*DB, error) {
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
	if err = db.AutoMigrate(&sctructs.User{}); err != nil {
		log.Println("AutoMigrate() -", err.Error())
	}
	if err = db.AutoMigrate(&sctructs.Score{}); err != nil {
		log.Println("AutoMigrate() -", err.Error())
	}
	if err = db.AutoMigrate(&sctructs.Item{}); err != nil {
		log.Println("AutoMigrate() -", err.Error())
	}
	return &DB{D: db}, nil
}

func InitItemsDB() (*DB, error) {
	creds := CredsForDB{
		Name:     "items",
		Host:     "localhost",
		Port:     "5432",
		User:     "maks",
		Password: "abc123",
	}

	db, err := gorm.Open(postgres.Open(creds.GetURI()), &gorm.Config{})
	if err != nil || db == nil {
		return nil, err
	}
	if err = db.AutoMigrate(&sctructs.Item{}); err != nil {
		log.Println("AutoMigrate() -", err.Error())
	}
	return &DB{D: db}, nil
}
