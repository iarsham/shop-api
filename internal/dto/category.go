package dto

type CategoryRequest struct {
	Title string `json:"title" binding:"required" example:"digital"`
}
