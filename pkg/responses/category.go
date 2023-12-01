package responses

import "github.com/iarsham/shop-api/internal/models"

type CategoryResponse models.Category

type CategoryExistsResponse struct {
	Response string `example:"this category already exists"`
}

type CategoryNotFoundResponse struct {
	Response string `example:"category not found"`
}

type CategoryDuplicateResponse struct {
	Response string `example:"category with this title already exists"`
}
