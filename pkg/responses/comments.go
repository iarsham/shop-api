package responses

import "github.com/iarsham/shop-api/internal/models"

type CommentResponse models.Comments

type CommentNotFoundResponse struct {
	Response string `example:"comment not found"`
}

type OwnerCantLikeCommentResponse struct {
	Response string `example:"owner can't like"`
}
