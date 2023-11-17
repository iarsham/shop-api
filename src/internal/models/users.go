package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FirstName string `gorm:"size:75;not null" json:"first_name"`
	LastName  string `gorm:"size:75;not null" json:"last_name"`
	Phone     string `gorm:"size:75;not null;index;unique" json:"phone"`
	Password  string `gorm:"size:300;not null" json:"-"`
	IsActive  bool   `gorm:"not null;default:false" json:"is_active"`
}
