package handlers

import (
	"go18/book/v3/controllers"

	"github.com/gin-gonic/gin"
)

var Comment = &CommentApiHandler{}

type CommentApiHandler struct {
}

func (h *CommentApiHandler) AddComment(ctx *gin.Context) {
	controllers.Book.GetBook(ctx, controllers.NewGetBookRequest(ctx.Param("bn")))
}
