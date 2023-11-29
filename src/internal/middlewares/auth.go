package middlewares

import (
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
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "authenticate required"})
			return
		}
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || strings.ToLower(authToken[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "authorization format is not correct"})
			return
		}
		claims, err := tokenService.GetClaims(authToken[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "token expired"})
				return
			default:
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "token invalid"})
				return
			}
		}
		if claims["sub"] != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "refresh token not allowed"})
			return
		}
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("phone", claims["phone"])
		ctx.Next()
	}
}

func IsAdminMiddleware(logs *common.Logger) gin.HandlerFunc {
	userService := services.NewUserService(logs)

	return func(ctx *gin.Context) {
		userID := ctx.GetString("user_id")
		if user, _ := userService.GetUserByID(userID); !user.IsAdmin {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"response": "permission not allowed, just admin user can perform this action"})
			return
		}
		ctx.Next()
	}
}
