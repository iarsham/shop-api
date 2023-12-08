package models

import "github.com/iarsham/shop-api/internal/common"

type Likes struct {
	common.ModelCreate
	CommentsID int `gorm:"column:comment_id"`
	UsersID    int `gorm:"column:users_id"`
}
