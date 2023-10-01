package dto

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required,min=1,max=75"`
	LastName  string `json:"last_name" binding:"required,min=1,max=75"`
	Phone     string `json:"phone" binding:"required,min=11,max=13"`
}
