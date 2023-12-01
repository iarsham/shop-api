package responses

type Token struct {
	Access  string `json:"access-token"`
	Refresh string `json:"refresh-token"`
}

type VerifyOTPResponse struct {
	Response Token
}

type OtpExpiredResponse struct {
	Response string `example:"otp expired"`
}

type SendOtpOkResponse struct {
	Response string `example:"otp was sent"`
}

type OtpIncorrectResponse struct {
	Response string `example:"code is incorrect"`
}
