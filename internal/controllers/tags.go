package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
)

type TagController struct {
	service *services.TagService
}

func NewTagController(logs *common.Logger) *TagController {
	return &TagController{
		service: services.NewTagService(logs),
	}
}

func (t *TagController) ListTagsHandler(ctx *gin.Context) {
	tags, err := t.service.GetAllTags()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, tags)
}

func (t *TagController) CreateTagHandler(ctx *gin.Context) {
	data := new(dto.TagRequest)
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, exists := t.service.GetTagByName(data.Name); exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{constans.Response: constans.TagByNameExists})
		return
	}
	createdTag, err := t.service.CreateTag(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusCreated, createdTag)
}

func (t *TagController) UpdateTagHandler(ctx *gin.Context) {
	data := new(dto.TagRequest)
	param := ctx.Param("pk")
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	if _, exists := t.service.GetTagByPK(param); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.TagNotFound})
		return
	}
	updatedTag, err := t.service.UpdateTag(param, data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, updatedTag)
}

func (t *TagController) DeleteTagHandler(ctx *gin.Context) {
	param := ctx.Param("pk")
	if _, exists := t.service.GetTagByPK(param); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.TagNotFound})
		return
	}
	if err := t.service.DeleteTag(param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
