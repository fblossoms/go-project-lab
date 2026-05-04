package impl

import (
	"context"
	"go18/book/v4/apps/book"
	"go18/book/v4/apps/book/impl"
	"go18/book/v4/apps/comment"
)

func (c CommentServiceImpl) AddComment(ctx context.Context, request *comment.AddCommentRequest) (*comment.Comment, error) {
	// 问题：能不能直接依赖Book.Service的具体实现？
	// 不能，强依赖具体业务实现，若后续技术栈升级，工程量激增
	// 应依赖接口，面向接口编程，不依赖业务实现
	//(&impl.BookServiceImpl{}).DescribeBook(ctx, nil)
	book.GetService().DescribeBook(ctx)
}
