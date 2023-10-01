package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
)

func SetupRoutes(r *gin.Engine, logs *common.Logger) {
	api := r.Group("/api/v1")

	userGroup := api.Group("/user")
	UsersRoutes(userGroup, logs)
}
