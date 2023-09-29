package configs

import (
	"github.com/iarsham/shop-api/internal/db"
	"log"
)

func InitialSrv() {
	err := db.OpenDB()
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	defer db.CloseDB()
	err = db.MigrateTables()
	if err != nil {
		log.Fatal("Tables migration failed: ", err)
	}
	log.Println("Postgres connected and Migration was successfully")
}
