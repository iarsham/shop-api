package services

import (
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentsService struct {
	db   *gorm.DB
	logs *common.Logger
}

func NewCommentsService(logs *common.Logger) *CommentsService {
	return &CommentsService{
		logs: logs,
		db:   db.GetDB(),
	}
}

func (c *CommentsService) CreateComment(req *dto.CommentsRequest, userID int, productSlug string) (*models.Comments, error) {
	var comment models.Comments
	comment.Message = req.Message
	comment.UsersID = userID
	comment.ProductsSlug = productSlug
	if err := c.db.Create(&comment).Error; err != nil {
		c.logs.Warn(err.Error())
		return nil, err
	}
	return &comment, nil
}

func (c *CommentsService) CommentByPK(pk string) (*models.Comments, bool) {
	var comment models.Comments
	err := c.db.Where("id=?", pk).First(&comment).Error
	if err != nil {
		c.logs.Warn(err.Error())
		return nil, false
	}
	return &comment, true
}

func (c *CommentsService) DeleteComment(pk string) error {
	comment, _ := c.CommentByPK(pk)
	if err := c.db.Select(clause.Associations).Delete(&comment).Error; err != nil {
		c.logs.Warn(err.Error())
		return err
	}
	return nil
}
