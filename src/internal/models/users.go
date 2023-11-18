package models

import "github.com/iarsham/shop-api/internal/common"

type Users struct {
	common.Model
	FirstName string     `gorm:"size:75;not null" json:"first_name"`
	LastName  string     `gorm:"size:75;not null" json:"last_name"`
	Phone     string     `gorm:"size:75;not null;index;unique" json:"phone"`
	Password  string     `gorm:"size:300;not null" json:"-"`
	IsActive  bool       `gorm:"not null;default:false" json:"is_active"`
	Comments  []Comments `gorm:"foreignKey:UsersID;references:ID"`
	Likes     []Likes    `gorm:"foreignKey:UsersID;references:ID"`
}
