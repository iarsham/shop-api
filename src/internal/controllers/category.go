package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"net/http"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(logs *common.Logger) *CategoryController {
	return &CategoryController{
		service: services.NewCategoryService(logs),
	}
}

func (c *CategoryController) GetCategoriesHandler(ctx *gin.Context) {
	categories, err := c.service.AllCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) CreateCategoryHandler(ctx *gin.Context) {
	data := new(dto.CategoryRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if _, exists := c.service.CategoryByPK(data.Title); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"response": "this category already exists"})
		return
	}
	createdCategory, err := c.service.CreateCategory(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusCreated, createdCategory)
}

func (c *CategoryController) UpdateCategoryHandler(ctx *gin.Context) {
	data := new(dto.CategoryRequest)
	pk := ctx.Param("pk")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}
	if _, exists := c.service.CategoryByPK(pk); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "category not found"})
		return
	}
	if _, exists := c.service.CategoryByTitle(data.Title); exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "category with this title already exists"})
		return
	}
	createdCategory, err := c.service.UpdateCategory(pk, data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, createdCategory)
}

func (c *CategoryController) DeleteCategoryHandler(ctx *gin.Context) {
	pk := ctx.Param("pk")
	if _, exists := c.service.CategoryByPK(pk); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "category not found"})
		return
	}
	if err := c.service.DeleteCategory(pk); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
