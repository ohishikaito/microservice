package domain

import (
	"time"
)

type User struct {
	Id              uint64 `gorm:"primary_key"`
	LastName        string
	FirstName       string
	Email           string
	TelephoneNumber string
	// 1 = 男性 2 = 女性
	Gender int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
