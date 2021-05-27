package db

import (
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/сfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitUsersDB() (*DB, error) {
	creds := CredsForDB{
		Name:     сfg.UserDB_name,
		Host:     сfg.UserDB_host,
		Port:     сfg.UserDB_port,
		User:     сfg.UserDB_user,
		Password: сfg.UserDB_password,
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
		Name:     сfg.ItemsDB_name,
		Host:     сfg.ItemsDB_host,
		Port:     сfg.ItemsDB_port,
		User:     сfg.ItemsDB_user,
		Password: сfg.ItemsDB_password,
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
