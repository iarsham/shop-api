package dto

import "mime/multipart"

type MediaRequest struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}
