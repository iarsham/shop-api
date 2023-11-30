package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"net/http"
	"strconv"
)

type OtpController struct {
	service      *services.OtpService
	serviceUser  *services.UserService
	serviceToken *services.TokenService
}

func NewOtpController(logs *common.Logger) *OtpController {
	service := services.NewOTPService(logs)
	return &OtpController{
		service:      service,
		serviceUser:  services.NewUserService(logs),
		serviceToken: services.NewTokenService(logs),
	}
}

// SendOTPHandler
//
//	@Summary		Send OTP
//	@Description	This endpoint receives the user's phone in request body and generates an otp. it then sends the otp to the user's phone via sms.
//	@Tags			OTP
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.SendOTPRequest				true	"send otp body"
//	@Success		200		{object}	responses.SendOtpOkResponse		"Success"
//	@Failure		404		{object}	responses.UserNotFoundResponse	"not found"
//	@Router			/otp/send [post]
func (o *OtpController) SendOTPHandler(ctx *gin.Context) {
	data := new(dto.SendOTPRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if _, exists := o.serviceUser.GetUserByPhone(data.Phone); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "user not found"})
		return
	}
	go o.service.SendOTP(data.Phone, ctx.Request)
	ctx.JSON(http.StatusOK, gin.H{"response": "otp was sent"})
}

// VerifyOTPHandler
//
//	@Summary		Verify OTP
//	@Description	this endpoint receives the user's otp code in request body.if code match, the verification is successfully.
//	@Tags			OTP
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.VerifyOTPRequest			true	"verify otp body"
//	@Success		200		{object}	responses.VerifyOTPResponse		"Success"
//	@Failure		410		{object}	responses.OtpExpiredResponse	"Expired"
//	@Failure		401		{object}	responses.OtpIncorrectResponse	"incorrect"
//	@Router			/otp/verify [post]
func (o *OtpController) VerifyOTPHandler(ctx *gin.Context) {
	data := new(dto.VerifyOTPRequest)
	user, _ := o.serviceUser.GetUserByIP(ctx.Request)
	userID := strconv.Itoa(int(user.ID))
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if expired := o.service.IsOtpExpire(ctx.Request); expired {
		ctx.AbortWithStatusJSON(http.StatusGone, gin.H{"response": "otp expired"})
		return
	}
	if ok := o.service.IsOtpEqual(data.Code, ctx.Request); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "code is incorrect"})
		return
	}
	if err := o.serviceUser.ActivateUser(ctx.Request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	access, _ := o.serviceToken.GenerateAccessToken(userID, user.Phone)
	refresh, _ := o.serviceToken.GenerateRefreshToken(userID)
	ctx.JSON(http.StatusOK, gin.H{"access-token": access, "refresh-token": refresh})
}
