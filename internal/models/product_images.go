package models

import "github.com/iarsham/shop-api/internal/common"

type ProductImages struct {
	common.Model
	URL          string
	ProductsSlug string `gorm:"column:product_slug"`
}
