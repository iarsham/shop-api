package models

import (
	"github.com/gosimple/slug"
	"github.com/iarsham/shop-api/internal/common"
	"gorm.io/gorm"
)

type Category struct {
	common.SlugModel
	Title    string     `gorm:"index;not null;unique" json:"title" example:"Mobile"`
	Products []Products `gorm:"foreignKey:CategorySlug;references:Slug"`
}

func (c *Category) BeforeCreate(*gorm.DB) error {
	c.Slug = slug.Make(c.Title)
	return nil
}
