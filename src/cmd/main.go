package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/configs"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/routers"
)

func main() {
	logs := common.NewLogger()
	configs.InitialSrv(logs)
	defer db.CloseDB(logs)
	r := gin.Default()
	routers.SetupRoutes(r, logs)
	err := r.Run(":8000")
	common.LogError(logs, err)
}
