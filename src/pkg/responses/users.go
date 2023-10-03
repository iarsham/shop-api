package responses

import "github.com/iarsham/shop-api/internal/dto"

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
