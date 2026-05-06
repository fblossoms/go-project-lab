package api

import (
	"go18/book/v4/apps/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
)

func (h *BookApiHandler) createBook(ctx *gin.Context) {
	req := book.NewCreateBookRequest()

	err := ctx.BindJSON(req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.CreateBook(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) queryBook(ctx *gin.Context) {
	// 参数处理
	req := book.NewQueryBookeRequest()
	req.Keywords = ctx.Query("keyword")

	pageNumber := ctx.Query("page_number")
	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			response.Failed(ctx, err)
			return
		}
		req.PageNumber = uint64(pnInt)
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		psInt, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			response.Failed(ctx, err)
			return
		}
		req.PageSize = uint64(psInt)
	}

	set, err := h.svc.QueryBook(ctx.Request.Context(), req)
	if err != nil {
		// 针对Response的统一封装
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, set)
}
