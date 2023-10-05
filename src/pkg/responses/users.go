package responses

import (
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
)

type RegisterOKResponse struct {
	Response string `example:"user created"`
}

type RegisterConflictResponse struct {
	Response string `example:"user with this phone already exists"`
}

type SendOtpOkResponse struct {
	Response string `example:"otp was sent"`
}

type UserNotFoundResponse struct {
	Response string `example:"user not found"`
}

type VerifyOTPResponse struct {
	dto.TokenDto
}

type OtpExpiredResponse struct {
	Response string `example:"otp expired"`
}

type OtpIncorrectResponse struct {
	Response string `example:"code is incorrect"`
}

type UserResponse struct {
	models.Users
}

type RefreshTokenResponse struct{
	Response string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5NDE4MTcsInBob25lIjoiKzk4OTAyMTMxMjIyNCIsInVzZXJfaWQiOiI1In0.hzmZdfltaMDWaiTwO8IG1uPEyXOsu3JBs6giU2BDeMI"`
}