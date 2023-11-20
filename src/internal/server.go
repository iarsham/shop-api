package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/middlewares"
	"github.com/iarsham/shop-api/internal/routers"
	"github.com/iarsham/shop-api/internal/validators"
)

func InitialSrv(logs *common.Logger) {
	validators.RegisterValidators(logs)

	err := db.OpenDB(logs)
	common.LogError(logs, err)

	err = db.MigrateTables(logs)
	common.LogError(logs, err)
	common.LogInfo(logs, "Postgres connected and Migration was successfully")

	err = db.RedisClient()
	common.LogError(logs, err)
	common.LogInfo(logs, "Redis connected successfully")

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
