package dto

type RegisterLoginRequest struct {
	FirstName string `json:"first_name" example:"James" binding:"omitempty,min=1,max=75"`
	LastName  string `json:"last_name" example:"Rodriguez" binding:"omitempty,min=1,max=75"`
	Phone     string `json:"phone" example:"+989021112299" binding:"required,phone,min=11,max=13"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" example:"James" binding:"required,min=1,max=75"`
	LastName  string `json:"last_name" example:"Rodriguez" binding:"required,min=1,max=75"`
}
