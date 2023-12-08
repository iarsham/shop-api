package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
	"github.com/iarsham/shop-api/internal/middlewares"
)

func ProductsRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewProductsController(logs)
	r.GET("/list", c.GetProductsHandler)
	r.POST("/:pk/create", middlewares.JwtAuthMiddleware(logs), middlewares.IsAdminMiddleware(logs), c.CreateProductHandler)
	r.PUT("/update/:pk", middlewares.JwtAuthMiddleware(logs), middlewares.IsAdminMiddleware(logs), c.UpdateProductHandler)
	r.DELETE("/delete/:pk", middlewares.JwtAuthMiddleware(logs), middlewares.IsAdminMiddleware(logs), c.DeleteProductHandler)
}
