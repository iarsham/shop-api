package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
)

type TokenController struct {
	service     *services.TokenService
	serviceUser *services.UserService
}

func NewTokenController(logs *common.Logger) *TokenController {
	service := services.NewTokenService(logs)
	serviceUser := services.NewUserService(logs)
	return &TokenController{
		service:     service,
		serviceUser: serviceUser,
	}
}

// RefreshTokenHandler
//
//	@Summary		Get New AccessToken
//	@Description	Create new access token from refresh token
//	@Tags			Tokens
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.RefreshTokenRequest				true	"refresh token body"
//	@Success		200		{object}	responses.RefreshTokenResponse		"Success"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/token/refresh-token [Post]
func (t *TokenController) RefreshTokenHandler(ctx *gin.Context) {
	data := new(dto.RefreshTokenRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, err := t.service.VerifyToken(data.RefreshToken); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	claims, err := t.service.GetClaims(data.RefreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	userID := claims["sub"].(string)
	user, _ := t.serviceUser.GetUserByID(userID)
	newAccessToken, _ := t.service.GenerateAccessToken(userID, user.Phone)
	ctx.JSON(http.StatusOK, gin.H{constans.Response: newAccessToken})
}
