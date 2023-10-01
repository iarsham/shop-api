package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FirstName string `gorm:"size:75;not null"`
	LastName  string `gorm:"size:75;not null"`
	Phone     string `gorm:"size:75;not null;index;unique"`
	Password  string `gorm:"size:300;not null"`
	IsActive  bool   `gorm:"not null;default:false"`
}
