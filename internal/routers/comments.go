package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/controllers"
)

func CommentsRoutes(r *gin.RouterGroup, logs *common.Logger) {
	c := controllers.NewCommentsController(logs)
	r.POST("/:pk/create", c.CreateCommentHandler)
	r.DELETE("/:pk/delete", c.DeleteCommentHandler)
}
