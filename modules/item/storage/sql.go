package storage

import (
	"gorm.io/gorm"
)

type SqlStorage struct {
	db *gorm.DB
}

func NewSQLStorage(db *gorm.DB) *SqlStorage {
	return &SqlStorage{
		db: db,
	}
}
