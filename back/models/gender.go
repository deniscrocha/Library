package models

type Gender struct {
	id          uint `gorm:"primaryKey"`
	name        string
	description string
}

func (Gender) TableName() string {
	return "Genders"
}
