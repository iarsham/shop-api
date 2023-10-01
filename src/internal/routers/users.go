package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
)

func UsersRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewUserController(logs)
	r.POST("/register", c.RegisterUserHandler)
}
