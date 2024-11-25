package models

type TypeUser string

const (
	BASIC  TypeUser = "BASIC"
	WORKER TypeUser = "WORKER"
	ADMIN  TypeUser = "ADMIN"
)

type User struct {
	id       uint `gorm:"primaryKey"`
	userType TypeUser
	name     string
	email    string
	password string
	enable   bool
}

func (User) TableName() string {
	return "Users"
}
