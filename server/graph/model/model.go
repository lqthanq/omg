package model

import "time"

type Model struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
