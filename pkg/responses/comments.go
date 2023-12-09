package responses

import "github.com/iarsham/shop-api/internal/models"

type CommentResponse models.Comments

type CommentNotFoundResponse struct {
	Response string `example:"comment not found"`
}
