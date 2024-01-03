package middlewares

import (
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/services"
)

func JwtAuthMiddleware(logs *common.Logger) gin.HandlerFunc {
	tokenService := services.NewTokenService(logs)

	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get(constans.Authorization)
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{constans.Response: constans.AuthenticateRequired})
			return
		}
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || strings.ToLower(authToken[0]) != constans.TokenType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{constans.Response: constans.AuthenticateFormat})
			return
		}
		claims, err := tokenService.GetClaims(authToken[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{constans.Response: constans.TokenExpired})
				return
			default:
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{constans.Response: constans.TokenInvalid})
				return
			}
		}
		if claims[constans.SUB] != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{constans.Response: constans.RefreshNotAllowed})
			return
		}
		ctx.Set(constans.UserID, claims[constans.UserID])
		ctx.Set(constans.Phone, constans.Phone)
		ctx.Next()
	}
}

func IsAdminMiddleware(logs *common.Logger) gin.HandlerFunc {
	userService := services.NewUserService(logs)

	return func(ctx *gin.Context) {
		userID := ctx.GetString(constans.UserID)
		if user, _ := userService.GetUserByID(userID); !user.IsAdmin {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{constans.Response: constans.PermissionAdminAllowed})
			return
		}
		ctx.Next()
	}
}
