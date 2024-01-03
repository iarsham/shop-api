package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
	"github.com/iarsham/shop-api/internal/middlewares"
)

func ProductImagesRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewProductImagesController(logs)
	r.POST(
		"/:pk/create",
		middlewares.JwtAuthMiddleware(logs),
		middlewares.IsAdminMiddleware(logs),
		c.CreateProductImageHandler,
	)
}
