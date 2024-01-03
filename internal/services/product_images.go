package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"github.com/iarsham/shop-api/pkg/constans"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type ProductImagesService struct {
	db   *gorm.DB
	logs *common.Logger
}

func NewProductImagesService(logs *common.Logger) *ProductImagesService {
	return &ProductImagesService{
		db:   db.GetDB(),
		logs: logs,
	}
}

func (p *ProductImagesService) CreateProductImages(req *dto.MediaRequest, param string, ctx *gin.Context) error {
	var images []models.ProductImages
	for _, file := range req.Files {
		url, exists := p.GetFileURL(file.Filename, ctx)
		if !exists {
			if err := ctx.SaveUploadedFile(file, constans.UploadPath+file.Filename); err != nil {
				p.logs.Warn(err.Error())
				return err
			} else {
				obj := models.ProductImages{
					ProductsSlug: param,
					URL:          url,
				}
				images = append(images, obj)
			}
		}
	}
	if images != nil {
		tx := p.db.Begin()
		if err := tx.Error; err != nil {
			p.logs.Fatal(err.Error())
			return err
		}
		if err := tx.Create(&images).Error; err != nil {
			p.logs.Warn(err.Error())
			return err
		}
		tx.Commit()
	}
	return nil
}

func (p *ProductImagesService) GetFileURL(fileName string, ctx *gin.Context) (string, bool) {
	scheme := "http://"
	if ctx.Request.TLS != nil {
		scheme = "https://"
	}
	baseURL := scheme + ctx.Request.Host
	filePath := filepath.Join("./uploads", fileName)
	fileURL := fmt.Sprintf("%s/media/%s", baseURL, fileName)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fileURL, false
	}
	return fileURL, true
}

func (p *ProductImagesService) ProductByPK(pk string) (*models.Products, bool) {
	var product models.Products
	err := p.db.Where("slug=?", pk).First(&product).Error
	if err != nil {
		p.logs.Warn(err.Error())
		return nil, false
	}
	return &product, true
}
