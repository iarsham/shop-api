package db

import (
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/models"
)

func MigrateTables(logs *common.Logger) error {
	var tables []interface{}
	db := GetDB()
	tables = append(tables, models.Users{})
	tables = append(tables, models.Category{})
	tables = append(tables, models.Products{})
	tables = append(tables, models.Comments{})
	tables = append(tables, models.Tags{})
	tables = append(tables, models.Likes{})
	err := db.Migrator().AutoMigrate(tables...)
	common.LogError(logs, err)
	return nil
}
