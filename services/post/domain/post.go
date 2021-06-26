package domain

import (
	"time"
)

type Post struct {
	Id     uint64 `gorm:"primary_key"`
	Text   string
	UserId int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
