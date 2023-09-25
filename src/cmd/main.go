package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "Hello From Gin"})
	})
	err := r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
