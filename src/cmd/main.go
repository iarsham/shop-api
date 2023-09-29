package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/configs"
	"log"
)

func main() {
	configs.InitialSrv()
	r := gin.Default()
	err := r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
