package models

import (
	"database/sql"
)

type BookModel struct {
	id           uint `gorm:"primaryKey"`
	name         string
	release_date sql.NullTime
	author       *string
	quantity     int
	series       *string
}

func (BookModel) TableName() string {
	return "BooksModels"
}
