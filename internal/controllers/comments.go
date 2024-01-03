package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/services"
	"github.com/iarsham/shop-api/pkg/constans"
	"net/http"
	"strconv"
)

type CommentsController struct {
	service        *services.CommentsService
	serviceProduct *services.ProductService
}

func NewCommentsController(logs *common.Logger) *CommentsController {
	return &CommentsController{
		service:        services.NewCommentsService(logs),
		serviceProduct: services.NewProductService(logs),
	}
}

// CreateCommentHandler
//
//	@Summary		Create comment for products
//	@Description	handler that is responsible for creating comment for products.
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		dto.CommentsRequest					true	"Create Comment Body"
//	@Param			pk		path		string								true	"Product Slug"
//	@Success		201		{object}	responses.CommentResponse			"Success"
//	@Failure		404		{object}	responses.ProductNOTExistsResponse	"Warn"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/comment/{pk}/create/ [post]
func (c *CommentsController) CreateCommentHandler(ctx *gin.Context) {
	data := new(dto.CommentsRequest)
	userID, _ := strconv.Atoi(ctx.GetString(constans.UserID))
	param := ctx.Param(constans.PK)
	if _, exists := c.serviceProduct.ProductByPK(param); !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.ProductNotFound})
		return
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{constans.Response: err.Error()})
		return
	}
	createdComment, err := c.service.CreateComment(data, userID, param)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusCreated, createdComment)
}

// DeleteCommentHandler
//
//	@Summary		Delete comment for products
//	@Description	handler that is responsible for deleting comment for products.
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			pk	path		string										true	"Comment ID"
//	@Success		204	{object}	responses.DeleteRecordResponse				"Success"
//	@Failure		404	{object}	responses.CommentNotFoundResponse			"Warn"
//	@Failure		403	{object}	responses.PermissionAdminAllowedResponse	"Warn"
//	@Failure		500	{object}	responses.InterServerErrorResponse			"Error"
//	@Router			/comment/{pk}/delete/ [delete]
func (c *CommentsController) DeleteCommentHandler(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString(constans.UserID))
	param := ctx.Param(constans.PK)
	comment, exists := c.service.CommentByPK(param)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.CommentNotFound})
		return
	}
	if userID != comment.UsersID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{constans.Response: constans.PermissionNotAllowed})
		return
	}
	if err := c.service.DeleteComment(param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// AddLikeToCommentHandler
//
//	@Summary		Add like to comment for products
//	@Description	handler that is responsible for add like to comment for products.
//	@Tags			Likes
//	@Accept			json
//	@Produce		json
//	@Param			pk	path		string									true	"Comment ID"
//	@Success		200	{object}	responses.Success						"Success"
//	@Failure		404	{object}	responses.CommentNotFoundResponse		"Warn"
//	@Failure		403	{object}	responses.OwnerCantLikeCommentResponse	"Warn"
//	@Failure		500	{object}	responses.InterServerErrorResponse		"Error"
//	@Router			/comment-likes/{pk}/add/ [post]
func (c *CommentsController) AddLikeToCommentHandler(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString(constans.UserID))
	param := ctx.Param(constans.PK)
	comment, exists := c.service.CommentByPK(param)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{constans.Response: constans.CommentNotFound})
		return
	}
	if userID == comment.UsersID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{constans.Response: constans.OwnerCantLike})
		return
	}
	if err := c.service.AddLike(param, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{constans.Response: constans.InternalServerResponse})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{constans.Response: constans.Success})
}
