package dto

type SendOTPRequest struct {
	Phone string `json:"phone" example:"+989021112299" binding:"required,phone,min=11,max=13"`
}

type VerifyOTPRequest struct {
	Code string `json:"code" example:"241960" binding:"required,min=6,max=6"`
}
