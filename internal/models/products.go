package models

import (
	"github.com/gosimple/slug"
	"github.com/iarsham/shop-api/internal/common"
	"gorm.io/gorm"
)

type Products struct {
	common.SlugModel
	Name         string     `gorm:"size:75;not null;index" example:"Phone"`
	Description  string     `gorm:"size:300;not null" example:"Phone Description"`
	Price        float64    `gorm:"not null" example:"599"`
	Stock        int        `gorm:"not null" example:"6"`
	IsAvailable  bool       `gorm:"not null" example:"true"`
	Weight       float64    `gorm:"not null" example:"0.7"`
	Views        int        `gorm:"not null"`
	CategorySlug string     `gorm:"not null" json:"category_slug" example:"digital"`
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
