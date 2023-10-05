package configs

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
)

func InitialSrv(logs *common.Logger) {
	RegisterValidators(logs)
	err := db.OpenDB(logs)
	common.LogError(logs, err)
	err = db.MigrateTables(logs)
	common.LogError(logs, err)
	logs.Info("Postgres connected and Migration was successfully")
	err = db.RedisClient()
	common.LogError(logs, err)
	logs.Info("Redis connected successfully")
}

func RegisterValidators(logs *common.Logger) {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("phone", common.IrPhoneValidator, true)
		common.LogError(logs, err)
	}
}
