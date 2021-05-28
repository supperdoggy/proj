package db

import (
	cfg2 "github.com/supperdoggy/score/cfg"
	"github.com/supperdoggy/score/sctructs"
	cfg "github.com/supperdoggy/score/сfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitUsersDB() (*DB, error) {
	creds := CredsForDB{
		Name:     cfg.UserDB_name,
		Host:     cfg.UserDB_host,
		Port:     cfg.UserDB_port,
		User:     cfg.UserDB_user,
		Password: cfg.UserDB_password,
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
		Name:     cfg.ItemsDB_name,
		Host:     cfg.ItemsDB_host,
		Port:     cfg.ItemsDB_port,
		User:     cfg.ItemsDB_user,
		Password: cfg2.ItemsDB_password,
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
