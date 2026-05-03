package impl

import (
	"context"
	"go18/book/v4/apps/book"

	"github.com/infraboard/mcube/v2/types"
)

func (b *BookServiceImpl) CreateBook(ctx context.Context, request *book.CreateBookRequest) (*book.Book, error) {
	panic("未知")
}

func (b *BookServiceImpl) QueryBook(ctx context.Context, request *book.QueryBookRequest) (*types.Set[*book.Book], error) {
	panic("未知")
}

func (b *BookServiceImpl) DescribeBook(ctx context.Context, request *book.DescribeBookRequest) (*book.Book, error) {
	panic("未知")
}

func (b *BookServiceImpl) UpdateBook(ctx context.Context, request *book.UpdateBookRequest) (*book.Book, error) {
	panic("未知")
}

func (b *BookServiceImpl) DeleteBook(ctx context.Context, request *book.DeleteBookRequest) (*book.Book, error) {
	panic("未知")
}
