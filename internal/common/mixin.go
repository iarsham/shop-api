package common

import (
	"time"
)

type Model struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type SlugModel struct {
	Slug      string    `gorm:"primaryKey" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type ModelCreate struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}
