package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
)

func MediaSizeMiddleware() gin.HandlerFunc {
	maxSize := int64(5 << 20)
	return func(ctx *gin.Context) {
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxSize)
		if ctx.Request.ContentLength > maxSize {
			ctx.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{constans.Response: constans.MaxFileSize})
			return
		}
		ctx.Next()
	}
}
