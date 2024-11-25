package models

type BookGender struct {
	id        uint      `gorm:"primaryKey"`
	bookModel BookModel `gorm:"foreignKey:id"`
	gender    Gender    `gorm:"foreignKey:id"`
}

func (BookGender) TableName() string {
	return "BooksGenders"
}
