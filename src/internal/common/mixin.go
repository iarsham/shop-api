package common

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type SlugModel struct {
	Slug      string         `gorm:"primaryKey" json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ModelCreate struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
