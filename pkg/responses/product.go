package responses

import "github.com/iarsham/shop-api/internal/models"

type ProductResponse models.Products

type ProductExistsResponse struct {
	Response string `example:"this product already exists"`
}

type ProductNOTExistsResponse struct {
	Response string `example:"product not found"`
}
