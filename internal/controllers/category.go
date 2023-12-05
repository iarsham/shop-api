package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"github.com/iarsham/shop-api/pkg/constans"
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

// GetCategoriesHandler
//
//	@Summary		Get All Categories
//	@Description	This endpoint returns a list of all categories in the store.
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.CategoryResponse			"Success"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/category/list [get]
func (c *CategoryController) GetCategoriesHandler(ctx *gin.Context) {
	categories, err := c.service.AllCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// CreateCategoryHandler
//
//	@Summary		Create New Category
//	@Description	This endpoint creates a new category in the store. The request body must contain the following information:
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.CategoryRequest					true	"create category body"
//	@Success		201		{object}	responses.CategoryResponse			"Success"
//	@Success		409		{object}	responses.CategoryExistsResponse	"Warn"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/category/create [post]
func (c *CategoryController) CreateCategoryHandler(ctx *gin.Context) {
	data := new(dto.CategoryRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, exists := c.service.CategoryByPK(data.Title); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{constans.Response: constans.CategoryExists})
		return
	}
	createdCategory, err := c.service.CreateCategory(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusCreated, createdCategory)
}

// UpdateCategoryHandler
//
//	@Summary		Update exists Category
//	@Description	This endpoint updates an existing Category in the store. The request body must contain the following information:
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.CategoryRequest					true	"update category body"
//	@Param			pk		path		string								true	"Category Slug"
//	@Success		200		{object}	responses.CategoryResponse			"Success"
//	@Failure		404		{object}	responses.CategoryNotFoundResponse	"Not Found"
//	@Failure		409		{object}	responses.CategoryDuplicateResponse	"Warn"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/category/update/{pk} [put]
func (c *CategoryController) UpdateCategoryHandler(ctx *gin.Context) {
	data := new(dto.CategoryRequest)
	pk := ctx.Param(constans.PK)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, exists := c.service.CategoryByPK(pk); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.CategoryNotFound})
		return
	}
	if _, exists := c.service.CategoryByTitle(data.Title); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{constans.Response: constans.CategoryByTitleExists})
		return
	}
	createdCategory, err := c.service.UpdateCategory(pk, data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, createdCategory)
}

// DeleteCategoryHandler
//
//	@Summary		Delete exists Category
//	@Description	This endpoint deletes an existing Category from the store.
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			pk	path		string								true	"Category Slug"
//	@Success		204	{object}	responses.DeleteRecordResponse		"Success"
//	@Failure		404	{object}	responses.CategoryNotFoundResponse	"Not Found"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/category/delete/{pk} [delete]
func (c *CategoryController) DeleteCategoryHandler(ctx *gin.Context) {
	pk := ctx.Param(constans.PK)
	if _, exists := c.service.CategoryByPK(pk); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.CategoryNotFound})
		return
	}
	if err := c.service.DeleteCategory(pk); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
