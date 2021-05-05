package db

import (
	"fmt"
	"gorm.io/gorm"
)

type CredsForDB struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

func (c *CredsForDB) GetURI() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.Name, c.Port)
}

type DB struct {
	D *gorm.DB
}

func (d DB) Close() (err error) {
	sqldb, err := d.D.DB()
	if err != nil || sqldb == nil {
		return
	}
	err = sqldb.Close()
	return
}

func (d DB) Create(value interface{}) error {
	err := d.D.Create(value).Error
	return err
}

func (d DB) AutoMigrate(a interface{}) error {
	return d.D.AutoMigrate(a)
}

func (d DB) Find(a interface{}) *gorm.DB {
	return d.D.Find(a)
}

func (d DB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return d.D.Where(query, args...)
}

func (d DB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return d.D.First(dest, conds...)
}
