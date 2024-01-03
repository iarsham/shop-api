package services

import (
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagService struct {
	db   *gorm.DB
	logs *common.Logger
}

func NewTagService(logs *common.Logger) *TagService {
	return &TagService{
		db:   db.GetDB(),
		logs: logs,
	}
}

func (t *TagService) CreateTag(req *dto.TagRequest) (*models.Tags, error) {
	var tag models.Tags
	tag.Name = req.Name
	if err := t.db.Save(&tag).Error; err != nil {
		t.logs.Warn(err.Error())
		return nil, err
	}
	return &tag, nil
}

func (t *TagService) UpdateTag(pk string, req *dto.TagRequest) (*models.Tags, error) {
	tag, _ := t.GetTagByPK(pk)
	if err := t.db.Model(&tag).Update("name", req.Name).Error; err != nil {
		t.logs.Warn(err.Error())
		return nil, err
	}
	return tag, nil
}

func (t *TagService) DeleteTag(pk string) error {
	tag, _ := t.GetTagByPK(pk)
	if err := t.db.Select(clause.Associations).Delete(&tag).Error; err != nil {
		t.logs.Warn(err.Error())
		return err
	}
	return nil
}

func (t *TagService) GetTagByPK(pk string) (*models.Tags, bool) {
	var tag models.Tags
	if err := t.db.Where("id=?", pk).First(&tag).Error; err != nil {
		t.logs.Warn(err.Error())
		return nil, false
	}
	return &tag, true
}

func (t *TagService) GetTagByName(name string) (*models.Tags, bool) {
	var tag models.Tags
	if err := t.db.Where("name=?", name).First(&tag).Error; err != nil {
		t.logs.Warn(err.Error())
		return nil, false
	}
	return &tag, true
}

func (t *TagService) GetAllTags() (*[]models.Tags, error) {
	var tags []models.Tags
	if err := t.db.Find(&tags).Error; err != nil {
		t.logs.Warn(err.Error())
		return nil, err
	}
	return &tags, nil
}
