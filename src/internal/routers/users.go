package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
	"github.com/iarsham/shop-api/internal/middlewares"
)

func UsersRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewUserController(logs)
	r.POST("/register", c.RegisterUserHandler)
	r.POST("/send-otp", c.SendOTPHandler)
	r.POST("/verify-otp", c.VerifyOTPHandler)
	r.GET("/", middlewares.JwtAuthMiddleware(logs), c.UserHandler)
	r.PUT("/", middlewares.JwtAuthMiddleware(logs), c.UserUpdateHandler)
}
