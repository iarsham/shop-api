package main

import (
	"github.com/gin-gonic/gin"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/configs"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/routers"
)

// @title			Shop API Document
// @contact.email   arshamdev2001@gmail.com
// @BasePath	 	/api/v1/
// @host			localhost:8000
// @licence.Name	MIT
// @licence.url     https://www.mit.edu/~amini/LICENSE.md
// @schemes			http https
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @accept json

func main() {
	r := gin.Default()
	logs := common.NewLogger()
	configs.InitialSrv(logs)
	defer db.CloseDB(logs)
	defer db.CloseRedis(logs)
	routers.SetupRoutes(r, logs)
	err := r.Run(":8000")
	common.LogError(logs, err)
}
