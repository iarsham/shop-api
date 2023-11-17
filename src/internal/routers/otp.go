package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
)

func OtpRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewOtpController(logs)
	r.POST("/send", c.SendOTPHandler)
	r.POST("/verify", c.VerifyOTPHandler)
}
