package models

import (
	"database/sql"
	"time"
)

type DeliveryStatus string

const (
	RENT      DeliveryStatus = "RENTED"
	LATE      DeliveryStatus = "LATE"
	DELIVERED DeliveryStatus = "DELIVERED"
)

type UserBook struct {
	id           uint `gorm:"primaryKey;"`
	userId       User `gorm:"foreignKey:id"`
	book         Book `gorm:"foreignKey:id"`
	pickupDate   time.Time
	expectedDate time.Time
	deliveryDate sql.NullTime
	price        *float32
	status       DeliveryStatus
}

func (UserBook) TableName() string {
	return "UsersBooks"
}
