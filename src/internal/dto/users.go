package dto

type RegisterRequest struct {
	FirstName string `json:"first_name" example:"James" binding:"required,min=1,max=75"`
	LastName  string `json:"last_name" example:"Rodriguez" binding:"required,min=1,max=75"`
	Phone     string `json:"phone" example:"+989021112299" binding:"required,min=11,max=13"`
}

type SendOTPRequest struct {
	Phone string `json:"phone" example:"+989021112299" binding:"required,min=11,max=13"`
}

type VerifyOTPRequest struct {
	Phone string `json:"phone" example:"+989021112299" binding:"required,min=11,max=13"`
	Code  string `json:"code" example:"241960" binding:"required,min=6,max=6"`
}
