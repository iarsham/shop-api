package services

import (
	"github.com/gosimple/slug"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductService struct {
	db    *gorm.DB
	redis *redis.Client
	logs  *common.Logger
}

func NewProductService(logs *common.Logger) *ProductService {
	return &ProductService{
		logs:  logs,
		db:    db.GetDB(),
		redis: db.GetRedis(),
	}
}

func (p *ProductService) AllProducts() (*[]models.Products, error) {
	var products []models.Products
	err := p.db.Preload("Comments").Preload("Tags").Find(&products).Error
	if err != nil {
		p.logs.Warn(err.Error())
		return nil, err
	}
	return &products, err
}

func (p *ProductService) CreateProduct(req *dto.ProductRequest) (*models.Products, error) {
	productObj, _ := common.TypeConverter[models.Products](req)
	err := p.db.Create(&productObj).Error
	if err != nil {
		p.logs.Warn(err.Error())
		return nil, err
	}
	return productObj, nil
}

func (p *ProductService) UpdateProduct(pk string, req *dto.ProductRequest) (*models.Products, error) {
	productObj, _ := p.ProductByPK(pk)
	productObj, _ = common.TypeConverter[models.Products](req)
	err := p.db.Model(&productObj).Update("slug", slug.Make(req.Name)).Error
	if err != nil {
		p.logs.Warn(err.Error())
		return nil, err
	}
	return productObj, nil
}

func (p *ProductService) DeleteProduct(pk string) error {
	product, _ := p.ProductByPK(pk)
	if err := p.db.Select(clause.Associations).Delete(&product).Error; err != nil {
		p.logs.Warn(err.Error())
		return err
	}
	return nil
}

func (p *ProductService) ProductByPK(pk string) (*models.Products, bool) {
	var product models.Products
	err := p.db.Where("slug=?", pk).First(&product).Error
	if err != nil {
		p.logs.Warn(err.Error())
		return nil, false
	}
	return &product, true
}

func (p *ProductService) ProductByName(name string) (*models.Products, bool) {
	var product models.Products
	err := p.db.Where("name=?", name).First(&product).Error
	if err != nil {
		p.logs.Warn(err.Error())
		return nil, false
	}
	return &product, true
}
