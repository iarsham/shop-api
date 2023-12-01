package responses

import (
	"github.com/iarsham/shop-api/internal/models"
)

type RegisterOKResponse struct {
	Response string `example:"Success, otp was sent"`
}

type LoginOKResponse struct {
	Response string `example:"Success, otp was sent"`
}

type UserNotFoundResponse struct {
	Response string `example:"user not found"`
}

type UserResponse models.Users
