package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/configs"
)

func main() {
	logs := common.NewLogger()
	configs.InitialSrv(logs)
	r := gin.Default()
	err := r.Run(":8000")
	common.LogError(logs, err)
}
