package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/api"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/middlewares"
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

	otpGroup := apiPrefix.Group("/otp")
	OtpRoutes(otpGroup, logs)

	tokenGroup := apiPrefix.Group("/token")
	TokenRoutes(tokenGroup, logs)

	categoryGroup := apiPrefix.Group("/category")
	CategoryRoutes(categoryGroup, logs)

	productsGroup := apiPrefix.Group("/product")
	ProductsRoutes(productsGroup, logs)

	productImagesGroup := apiPrefix.Group("/product-images")
	productImagesGroup.Use(middlewares.MediaSizeMiddleware())
	ProductImagesRoutes(productImagesGroup, logs)

	api.SwaggerInfo.BasePath = BaseURL
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

}
