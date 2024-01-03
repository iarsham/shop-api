package dto

type TagRequest struct {
	Name string `json:"name" example:"ai" binding:"required"`
}
