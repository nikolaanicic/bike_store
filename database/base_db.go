package database

import "gorm.io/gorm"

type BaseDatabase struct {
	db *gorm.DB
}

func NewBaseDB(db *gorm.DB) *BaseDatabase {
	return &BaseDatabase{db: db}
}
