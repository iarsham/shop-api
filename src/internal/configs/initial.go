package configs

import (
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
)

func InitialSrv(logs *common.Logger) {
	err := db.OpenDB(logs)
	common.LogError(logs, err)
	err = db.MigrateTables(logs)
	common.LogError(logs, err)
	logs.Info("Postgres connected and Migration was successfully")
	err = db.RedisClient()
	common.LogError(logs, err)
	logs.Info("Redis connected successfully")
}
