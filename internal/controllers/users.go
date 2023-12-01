package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
)

type UserController struct {
	service      *services.UserService
	serviceOtp   *services.OtpService
	serviceToken *services.TokenService
}

func NewUserController(logs *common.Logger) *UserController {
	service := services.NewUserService(logs)
	serviceOtp := services.NewOTPService(logs)
	serviceToken := services.NewTokenService(logs)
	return &UserController{
		service:      service,
		serviceOtp:   serviceOtp,
		serviceToken: serviceToken,
	}
}

// RegisterLoginUserHandler
//
//	@Summary		Register And Login By Phone
//	@Description	Create user with firstname / lastname / phone
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.RegisterLoginRequest			true	"register and login body"
//	@Success		201		{object}	responses.RegisterOKResponse		"Success"
//	@Success		200		{object}	responses.LoginOKResponse			"Success"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/user/register-login [post]
func (u *UserController) RegisterLoginUserHandler(ctx *gin.Context) {
	data := new(dto.RegisterLoginRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if _, exists := u.service.GetUserByPhone(data.Phone); exists {
		go u.serviceOtp.SendOTP(data.Phone, ctx.Request)
		ctx.JSON(http.StatusOK, gin.H{"response": "Success, otp was sent"})
		return
	}
	if err := u.service.RegisterLoginByPhone(data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	go u.serviceOtp.SendOTP(data.Phone, ctx.Request)
	ctx.JSON(http.StatusCreated, gin.H{"response": "Success, otp was sent"})
}

// UserHandler
//
//	@Summary		Get User
//	@Description	Retrieve user information by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.UserResponse				"Success"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/user/ [Get]
func (u *UserController) UserHandler(ctx *gin.Context) {
	id := ctx.GetString("user_id")
	userData, exists := u.service.GetUserByID(id)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, userData)
}

// UserUpdateHandler
//
//	@Summary		Update User
//	@Description	Update user information by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.UpdateUserRequest				true	"update user body"
//	@Success		200		{object}	responses.UserResponse				"Success"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/user/ [Put]
func (u *UserController) UserUpdateHandler(ctx *gin.Context) {
	id := ctx.GetString("user_id")
	data := new(dto.UpdateUserRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	userData, err := u.service.UpdateUserByID(id, data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, userData)
}
