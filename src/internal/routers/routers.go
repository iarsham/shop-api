package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/api"
	"github.com/iarsham/shop-api/internal/common"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RedirectToDocs(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}

func SetupRoutes(r *gin.Engine, logs *common.Logger) {
	r.GET("/", RedirectToDocs)
	BaseURL := "/api/v1"
	apiPrefix := r.Group(BaseURL)

	userGroup := apiPrefix.Group("/user")
	UsersRoutes(userGroup, logs)

	api.SwaggerInfo.BasePath = BaseURL
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

}
