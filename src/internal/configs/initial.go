package configs

import (
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
)

func InitialSrv(logs *common.Logger) {
	err := db.OpenDB(logs)
	common.LogError(logs, err)
	defer db.CloseDB(logs)

	err = db.MigrateTables(logs)
	common.LogError(logs, err)

	logs.Info("Postgres connected and Migration was successfully")
}
