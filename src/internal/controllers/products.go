package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
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

func (p *ProductsController) GetProductsHandler(ctx *gin.Context) {
	products, err := p.service.AllProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductsController) CreateProductHandler(ctx *gin.Context) {
	data := new(dto.ProductRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if _, exists := p.service.ProductByName(data.Name); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"response": "this product already exists"})
		return
	}
	createdUser, err := p.service.CreateProduct(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusCreated, createdUser)
}

func (p *ProductsController) UpdateProductHandler(ctx *gin.Context) {
	data := new(dto.ProductRequest)
	pk := ctx.Param("pk")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if _, exists := p.service.ProductByPK(pk); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "category not found"})
		return
	}
	updatedProduct, err := p.service.UpdateProduct(pk, data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductsController) DeleteProductHandler(ctx *gin.Context) {
	param := ctx.Param("pk")
	if _, exists := p.service.ProductByPK(param); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "product not found"})
		return
	}
	if err := p.service.DeleteProduct(param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
