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

// SendOTPHandler
// @Summary Send OTP
// @Description This endpoint receives the user's phone in request body and generates an otp. it then sends the otp to the user's phone via sms.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.SendOTPRequest true "send otp body"
// @Success 200 {object} responses.SendOtpOkResponse "Success"
// @Success 404 {object} responses.UserNotFoundResponse "not found"
// @Router /user/send-otp [post]
func (u *UserController) SendOTPHandler(ctx *gin.Context) {
	data := new(dto.SendOTPRequest)

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if exists := u.service.UserExistsByPhone(data.Phone); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "user not found"})
		return
	}

	go u.service.SendOTP(data.Phone)

	ctx.JSON(http.StatusOK, gin.H{"response": "otp was sent"})
}

// VerifyOTPHandler
// @Summary Verify OTP
// @Description this endpoint receives the user's phone and otp code in request body.if code match, the verification is successfully.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.VerifyOTPRequest true "verify otp body"
// @Success 200 {object} responses.VerifyOTPResponse "Success"
// @Success 410 {object} responses.OtpExpiredResponse "Expired"
// @Success 401 {object} responses.OtpIncorrectResponse "incorrect"
// @Router /user/verify-otp [post]
func (u *UserController) VerifyOTPHandler(ctx *gin.Context) {
	data := new(dto.VerifyOTPRequest)

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if isExpired := u.service.CheckOtpExpire(data.Phone); isExpired {
		ctx.AbortWithStatusJSON(http.StatusGone, gin.H{"response": "otp expired"})
		return
	}

	if ok := u.service.CheckOtpEqual(data.Phone, data.Code); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "code is incorrect"})
		return
	}

	token, err := u.service.VerifyUser(data.Phone)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
