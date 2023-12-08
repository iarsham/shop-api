package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
)

type ProductImagesController struct {
	service *services.ProductImagesService
}

func NewProductImagesController(logs *common.Logger) *ProductImagesController {
	return &ProductImagesController{
		service: services.NewProductImagesService(logs),
	}
}

func (p *ProductImagesController) CreateProductImageHandler(ctx *gin.Context) {
	data := new(dto.MediaRequest)
	param := ctx.Param(constans.PK)
	if _, exists := p.service.ProductByPK(param); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.ProductNotFound})
		return
	}
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if err := p.service.CreateProductImages(data, param, ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{constans.Response: constans.Created})
}
