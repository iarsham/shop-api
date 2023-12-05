package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
)

type ProductsController struct {
	service *services.ProductService
}

func NewProductsController(logs *common.Logger) *ProductsController {
	return &ProductsController{
		service: services.NewProductService(logs),
	}
}

// GetProductsHandler
//
//	@Summary		Get All Products
//	@Description	This endpoint returns a list of all Products in the store.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.ProductResponse			"Success"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/product/list [get]
func (p *ProductsController) GetProductsHandler(ctx *gin.Context) {
	products, err := p.service.AllProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// CreateProductHandler
//
//	@Summary		Create New Product
//	@Description	Creates a new product record in the database.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.ProductRequest					true	"create product body"
//	@Success		201		{object}	responses.ProductResponse			"Success"
//	@Success		409		{object}	responses.ProductExistsResponse		"Warn"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/product/create [post]
func (p *ProductsController) CreateProductHandler(ctx *gin.Context) {
	data := new(dto.ProductRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, exists := p.service.ProductByName(data.Name); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{constans.Response: constans.ProductExists})
		return
	}
	createdUser, err := p.service.CreateProduct(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusCreated, createdUser)
}

// UpdateProductHandler
//
//	@Summary		Update Exists Product
//	@Description	Update an exists product record in the database.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			pk		path		string								true	"Product Slug"
//	@Param			Request	body		dto.ProductRequest					true	"update product body"
//	@Success		200		{object}	responses.ProductResponse			"Success"
//	@Success		409		{object}	responses.ProductNOTExistsResponse	"Warn"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/product/update [put]
func (p *ProductsController) UpdateProductHandler(ctx *gin.Context) {
	data := new(dto.ProductRequest)
	pk := ctx.Param(constans.PK)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, exists := p.service.ProductByPK(pk); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.ProductNotFound})
		return
	}
	updatedProduct, err := p.service.UpdateProduct(pk, data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, updatedProduct)
}

// DeleteProductHandler
//
//	@Summary		Delete exists Product
//	@Description	This endpoint deletes an existing Product from the store.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			pk	path		string								true	"Product Slug"
//	@Success		204	{object}	responses.DeleteRecordResponse		"Success"
//	@Failure		404	{object}	responses.ProductNOTExistsResponse	"Not Found"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/product/delete/{pk} [delete]
func (p *ProductsController) DeleteProductHandler(ctx *gin.Context) {
	param := ctx.Param(constans.PK)
	if _, exists := p.service.ProductByPK(param); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.ProductNotFound})
		return
	}
	if err := p.service.DeleteProduct(param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
