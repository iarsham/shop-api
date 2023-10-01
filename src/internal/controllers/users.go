package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"net/http"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(logs *common.Logger) *UserController {
	service := services.NewUserService(logs)
	return &UserController{service: service}
}

func (u *UserController) RegisterUserHandler(ctx *gin.Context) {
	data := new(dto.RegisterRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if err := u.service.RegisterByPhone(data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"response": "user created"})
}
