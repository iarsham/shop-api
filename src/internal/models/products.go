package models

import (
	"github.com/gosimple/slug"
	"github.com/iarsham/shop-api/internal/common"
	"gorm.io/gorm"
)

type Products struct {
	common.SlugModel
	Name         string     `gorm:"size:75;not null;index"`
	Description  string     `gorm:"size:300;not null"`
	Price        float64    `gorm:"not null"`
	Stock        int        `gorm:"not null"`
	IsAvailable  bool       `gorm:"not null"`
	Weight       int        `gorm:"not null"`
	Views        int        `gorm:"not null"`
	CategorySlug string     `gorm:"not null" json:"category_slug"`
	Comments     []Comments `gorm:"foreignKey:ProductsSlug;references:Slug"`
	Tags         []Tags     `gorm:"many2many:product_tags"`
}

func (p *Products) BeforeSave(*gorm.DB) error {
	p.Slug = slug.Make(p.Name)
	if p.Stock > 0 {
		p.IsAvailable = true
	} else {
		p.IsAvailable = false
	}
	return nil
}
