package models

import "github.com/iarsham/shop-api/internal/common"

type Likes struct {
	common.ModelCreate
	CommentsID int `gorm:"comment_id"`
	UsersID    int `gorm:"users_id"`
}
