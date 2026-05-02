package controllers

import (
	"context"
	"fmt"
	"go18/book/v3/models"
)

var Comment = &CommentController{}

type CommentController struct {
}

type AddCommentRequest struct {
	BookNumber string
}

func (c *CommentController) AddComment(ctx context.Context, in AddCommentRequest) (*models.Comment, error) {
	book, err := Book.GetBook(ctx, NewGetBookRequest(in.BookNumber))

	// 异常断言
	//if exception.IsApiException(err, exception.CODE_NOT_FOUND) {
	//
	//}
	if err != nil {
		// 获取的Book是否存在（获取时是否报错）
		return nil, err
	}

	// 判断Book的状态
	fmt.Println(book)
	return nil, nil
}
