package domain

import "time"

type PostDetail struct {
	Id          uint64 `gorm:"primary_key"`
	Description string
	PostId      uint64

	CreatedAt time.Time
	UpdatedAt time.Time
}
