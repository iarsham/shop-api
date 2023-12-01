package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5NDE4MTcsInBob25lIjoiKzk4OTAyMTMxMjIyNCIsInVzZXJfaWQiOiI1In0.hzmZdfltaMDWaiTwO8IG1uPEyXOsu3JBs6giU2BDeMI"`
}
