package models

import "github.com/iarsham/shop-api/internal/common"

type Tags struct {
	common.ModelCreate
	Name string `gorm:"not null;index;unique"`
}
