package dto

type TokenDto struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzMzg4MTcsInBob25lIjoiKzk4OTAyMTMxMjIyNCIsInVzZXJfaWQiOiI1In0.DAxOeyiWpPZWXyVnnyLajMQ9SGsBKw65qOAurhjlFy0"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5NDE4MTcsInBob25lIjoiKzk4OTAyMTMxMjIyNCIsInVzZXJfaWQiOiI1In0.hzmZdfltaMDWaiTwO8IG1uPEyXOsu3JBs6giU2BDeMI"`
}
