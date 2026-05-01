package handlers

import (
	"go18/book/v3/config"
	"go18/book/v3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Book = &BookApiHandler{}

type BookApiHandler struct{}

func (h *BookApiHandler) Registry(server gin.IRouter) {
	server.GET("/api/books", Book.ListBook)

	server.POST("/api/books", Book.CreateBook)

	server.GET("/api/books/:bn", Book.GetBook)

	server.PUT("api/books/:bn", Book.UpdateBook)

	server.DELETE("api/books/:bn", Book.DeleteBook)
}

func (h *BookApiHandler) ListBook(ctx *gin.Context) {
	set := &models.BookSet{}

	pn, ps := 1, 20
	pageNumber := ctx.Query("page_number")
	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		pn = int(pnInt)
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		psInt, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		ps = int(psInt)
	}

	query := config.DB().Model(&models.Book{})

	kws := ctx.Query("Keywords")
	if kws != "" {
		query = query.Where("title LIKE ?", "%"+kws+"%")
	}

	offset := (pn - 1) * ps
	err := query.Count(&set.Total).Offset(int(offset)).Limit(int(ps)).Find(&set.Items).Error
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}
	ctx.JSON(200, set)
}

func (h *BookApiHandler) CreateBook(ctx *gin.Context) {
	bookSpecInstance := &models.BookSpec{}

	err := ctx.BindJSON(bookSpecInstance)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	bookInstance := &models.Book{BookSpec: *bookSpecInstance}

	err = config.DB().Save(bookInstance).Error
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, bookInstance)
}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	bookInstance := &models.Book{}

	err := config.DB().Where("id = ?", ctx.Param("bn")).Take(bookInstance).Error
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	bnStr := ctx.Param("bn")
	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	bookInstance := models.Book{
		Id: uint(bn),
	}

	err = ctx.BindJSON(&bookInstance.BookSpec)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	err = config.DB().Where("id = ?", bookInstance.Id).Updates(bookInstance).Error
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	err := config.DB().Where("id = ?", ctx.Param("bn")).Delete(&models.Book{}).Error
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "ok")
}
