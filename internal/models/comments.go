package models

import "github.com/iarsham/shop-api/internal/common"

type Comments struct {
	common.Model
	Message      string `gorm:"not null;index"`
	UsersID      int    `gorm:"column:user_id"`
	ProductsSlug string `gorm:"column:product_slug"`
	Likes        uint
}

type CommentLikes struct {
	common.Model
	CommentsID int `gorm:"column:comment_id"`
	UsersID    int `gorm:"column:user_id"`
}
