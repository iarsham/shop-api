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

// CreateProductImageHandler
//
//	@Summary		Create Image for products
//	@Description	handler that is responsible for creating product images.
//	@Tags			Product Images
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.MediaRequest						true	"Create Product Image Body"
//	@Param			pk		path		string									true	"Product Slug"
//	@Success		201		{object}	responses.CreateProductImagesResponse	"Success"
//	@Failure		409		{object}	responses.ProductNOTExistsResponse		"Warn"
//	@Failure		500		{object}	responses.InterServerErrorResponse		"Error"
//	@Router			/product-images/{pk}/create/ [post]
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
