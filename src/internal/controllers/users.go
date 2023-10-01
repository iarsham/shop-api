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

// RegisterUserHandler
// @Summary Register By Phone
// @Description Create user with firstname / lastname / phone
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterRequest true "register body"
// @Success 201 {object} responses.RegisterOKResponse "Success"
// @Success 409 {object} responses.RegisterConflictResponse "Conflict"
// @Success 500 {object} responses.InterServerErrorResponse "Error"
// @Router /user/register [post]
func (u *UserController) RegisterUserHandler(ctx *gin.Context) {
	data := new(dto.RegisterRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if exists := u.service.UserExistsByPhone(data.Phone); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"response": "user with this phone already exists"})
		return
	}

	if err := u.service.RegisterByPhone(data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"response": "user created"})
}
