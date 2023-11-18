package models

import "github.com/iarsham/shop-api/internal/common"

type Products struct {
	common.SlugModel
	Name         string     `gorm:"not null;index"`
	Description  string     `gorm:"not null"`
	Price        float64    `gorm:"not null"`
	Stock        int        `gorm:"not null"`
	IsAvailable  bool       `gorm:"not null"`
	Weight       int        `gorm:"not null"`
	Views        int        `gorm:"not null"`
	CategorySlug string     `gorm:"not null"`
	Comments     []Comments `gorm:"foreignKey:ProductsSlug;references:Slug"`
	Tags         []Tags     `gorm:"many2many:product_tags"`
}
