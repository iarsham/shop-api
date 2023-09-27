package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	HOST     = os.Getenv("PG_HOST")
	USER     = os.Getenv("PG_USER")
	PASSWORD = os.Getenv("PG_PASS")
	DB       = os.Getenv("PG_DB")
	PORT     = os.Getenv("PG_PORT")
	PgDB     *gorm.DB
)

func OpenDB() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DB, PORT)
	PgDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db, _ := PgDB.DB()
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return PgDB
}

func CloseDB() {
	db, _ := PgDB.DB()
	err := db.Close()
	if err != nil {
		return
	}
}
