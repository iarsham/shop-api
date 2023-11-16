package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/middlewares"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
)

func UsersRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewUserController(logs)
	r.POST("/register-login", c.RegisterLoginUserHandler)
	r.POST("/send-otp", c.SendOTPHandler)
	r.POST("/verify-otp", c.VerifyOTPHandler)
	r.POST("/refresh-token", c.RefreshTokenHandler)
	r.GET("/", middlewares.JwtAuthMiddleware(logs), c.UserHandler)
	r.PUT("/", middlewares.JwtAuthMiddleware(logs), c.UserUpdateHandler)
}
