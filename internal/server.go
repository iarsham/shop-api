package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/middlewares"
	"github.com/iarsham/shop-api/internal/routers"
	"github.com/iarsham/shop-api/internal/validators"
	"os"
)

func InitialSrv(logs *common.Logger) {
	validators.RegisterValidators(logs)

	err := db.OpenDB(logs)
	common.LogError(logs, err)
	logs.Info("Postgres connected successfully")

	err = db.RedisClient()
	common.LogError(logs, err)
	common.LogInfo(logs, "Redis connected successfully")

	if args := os.Args; len(args) > 1 {
		if args[1] == "createadmin" {
			db.CreateAdminUser(logs)
		}
		if args[1] == "migrate" {
			db.MigrateTables(logs)
		}
	}

	r := gin.Default()
	routers.SetupRoutes(r, logs)
	r.Use(middlewares.CorsMiddleware())
	err = r.Run()
	common.LogError(logs, err)
}

func CloseDB(logs *common.Logger) {
	db.CloseDB(logs)
	db.CloseRedis(logs)
}
