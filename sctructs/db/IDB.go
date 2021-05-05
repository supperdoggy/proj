package db

import "gorm.io/gorm"

type IDB interface {
	Create(value interface{}) error
	Find(a interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
}
