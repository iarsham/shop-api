package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
)

func TokenRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewTokenController(logs)
	r.POST("/refresh", c.RefreshTokenHandler)
}
