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
	Gender          int

	CreatedAt time.Time
	UpdatedAt time.Time
}
