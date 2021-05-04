package handlers

import "gorm.io/gorm"

type IDB interface {
	Create(value interface{}) error
	Find(a interface{}) error
	Where(query interface{}, args ...interface{}) *gorm.DB
}
