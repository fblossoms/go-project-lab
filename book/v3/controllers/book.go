package controllers

import (
	"context"
	"go18/book/v3/config"
	"go18/book/v3/models"
)

var Book = &BookController{}

type BookController struct {
}

type GetBookRequest struct {
	BookNumber string
}

// NewGetBookRequest 增加可拓展性
func NewGetBookRequest(bookNumber string) *GetBookRequest {
	return &GetBookRequest{
		BookNumber: bookNumber,
	}
}

// GetBook 核心功能
// ctx: Trace，支持请求的取消，request_id
// GetBookRequest: 为什么要把他封装为1个对象？GetBook(ctx context.Context, BookNumber string)保证了接口签名的兼容性
func (c *BookController) GetBook(ctx context.Context, in *GetBookRequest) (*models.Book, error) {
	//context.WithValue(ctx, "request_id", 111) // 请求上下文
	//ctx.Value("request_id") // 使用完后获取id

	config.L().Debug().Msgf("get book %s", in.BookNumber)

	bookInstance := &models.Book{}

	err := config.DB().Where("id = ?", in.BookNumber).Take(bookInstance).Error
	if err != nil {
		return nil, err
	}

	return bookInstance, nil
}

func (c *BookController) CreateBook(ctx context.Context, in *models.BookSpec) (*models.Book, error) {
	bookInstance := &models.Book{BookSpec: *in}

	err := config.DB().Save(bookInstance).Error
	if err != nil {
		return nil, err
	}
	return bookInstance, nil
}
