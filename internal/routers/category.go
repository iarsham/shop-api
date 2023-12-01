package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
	"github.com/iarsham/shop-api/internal/middlewares"
)

func CategoryRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewCategoryController(logs)
	r.GET("/list", c.GetCategoriesHandler)
	r.POST("/create", middlewares.JwtAuthMiddleware(logs), middlewares.IsAdminMiddleware(logs), c.CreateCategoryHandler)
	r.PUT("/update/:pk", middlewares.JwtAuthMiddleware(logs), middlewares.IsAdminMiddleware(logs), c.UpdateCategoryHandler)
	r.DELETE("/delete/:pk", middlewares.JwtAuthMiddleware(logs), middlewares.IsAdminMiddleware(logs), c.DeleteCategoryHandler)
}
