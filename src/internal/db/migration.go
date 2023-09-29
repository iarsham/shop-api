package db

import "github.com/iarsham/shop-api/internal/models"

func MigrateTables() error {
	var tables []interface{}
	db := GetDB()
	tables = append(tables, models.Users{})
	if err := db.Migrator().AutoMigrate(tables...); err != nil {
		return err
	}
	return nil
}
