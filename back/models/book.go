package models

type BookStatus string

const (
	RENTED    BookStatus = "RENTED"
	AVAILABLE BookStatus = "AVAILABLE"
	DAMAGED   BookStatus = "DAMAGED"
)

type Book struct {
	id        uint       `gorm:"primaryKey"`
	bookModel BookModel  `gorm:"foreignKey:id"`
	status    BookStatus `gorm:"type:enum('RENTED', 'AVAILABLE', 'DAMAGED')" gorm:"column:status"`
}

func (Book) TableName() string {
	return "Books"
}
