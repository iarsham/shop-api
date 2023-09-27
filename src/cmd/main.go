package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/db"
	"log"
)

func main() {
	err := db.OpenDB()
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	defer db.CloseDB()

	r := gin.Default()
	err = r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
