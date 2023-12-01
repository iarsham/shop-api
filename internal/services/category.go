package services

import (
	"github.com/gosimple/slug"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryService struct {
	db   *gorm.DB
	logs *common.Logger
}

func NewCategoryService(logs *common.Logger) *CategoryService {
	return &CategoryService{
		logs: logs,
		db:   db.GetDB(),
	}
}

func (c *CategoryService) AllCategories() (*[]models.Category, error) {
	var categories []models.Category
	err := c.db.Preload("Products").Find(&categories).Error
	if err != nil {
		c.logs.Warn(err.Error())
		return nil, err
	}
	return &categories, err
}

func (c *CategoryService) CreateCategory(req *dto.CategoryRequest) (*models.Category, error) {
	var category models.Category
	category.Title = req.Title
	err := c.db.Create(&category).Error
	if err != nil {
		c.logs.Warn(err.Error())
		return nil, err
	}
	return &category, nil
}

func (c *CategoryService) UpdateCategory(pk string, req *dto.CategoryRequest) (*models.Category, error) {
	categoryObj, _ := c.CategoryByPK(pk)
	newPK := slug.Make(req.Title)
	if err := c.db.Model(categoryObj).Updates(map[string]interface{}{"title": req.Title, "slug": newPK}).Error; err != nil {
		return nil, err
	}
	return categoryObj, nil
}

func (c *CategoryService) DeleteCategory(pk string) error {
	category, _ := c.CategoryByPK(pk)
	if err := c.db.Select(clause.Associations).Delete(&category).Error; err != nil {
		c.logs.Warn(err.Error())
	}
	return nil
}

func (c *CategoryService) CategoryByPK(pk string) (*models.Category, bool) {
	var category models.Category
	err := c.db.Where("slug=?", pk).First(&category).Error
	if err != nil {
		c.logs.Warn(err.Error())
		return nil, false
	}
	return &category, true
}

func (c *CategoryService) CategoryByTitle(title string) (*models.Category, bool) {
	var category models.Category
	err := c.db.Where("title=?", title).First(&category).Error
	if err != nil {
		c.logs.Warn(err.Error())
		return nil, false
	}
	return &category, true
}
