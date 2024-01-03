package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
)

func TagsRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewTagController(logs)
	r.GET("/list", c.ListTagsHandler)
	r.POST("/create", c.CreateTagHandler)
	r.PUT("/:pk/update", c.UpdateTagHandler)
	r.DELETE("/:pk/delete", c.DeleteTagHandler)
}
