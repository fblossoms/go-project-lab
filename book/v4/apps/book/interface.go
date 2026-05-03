package book

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/types"
)

// Service Book的业务定义
type Service interface {
	// CreateBook 1. 创建（录入）书籍
	CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error)

	// QueryBook 2. Book列表查询
	QueryBook(ctx context.Context, request *QueryBookRequest) (*types.Set[*Book], error)

	// DescribeBook 3. Book详情查询
	DescribeBook(ctx context.Context, request *DescribeBookRequest) (*Book, error)

	// UpdateBook 4. Book更新
	UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error)

	// DeleteBook  5. Book删除
	DeleteBook(ctx context.Context, request *DeleteBookRequest) (*Book, error)
}

type BookSet struct {
	// 总共有多少个
	Total int64 `json:"total"`
	// book清单
	Itmes []*Book `json:"itmes"`
}

type CreateBookRequest struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price" validate:"required"`
	IsSale bool    `json:"is_sale" gorm:"column:is_sale"`
}

func NewQueryBookeRequest() *QueryBookRequest {
	return &QueryBookRequest{
		// PageRequest{PageSize: 20, PageNumber: 1}
		PageRequest: *request.NewDefaultPageRequest(),
	}
}

type QueryBookRequest struct {
	// 下面两个属性可能涉及多类使用
	//PageSize   uint
	//PageNumber uint
	request.PageRequest // mcube已经封装，可以直接复用

	// 关键参数
	Keywords string `json:"keywords"`
}

type DescribeBookRequest struct {
	// 详细信息，如书籍对应的ISBN码
	Id uint
}

type UpdateBookRequest struct {
	// 找到要更新的书籍要使用ISBN码
	DescribeBookRequest

	// 更新时相当于更改信息
	CreateBookRequest
}

type DeleteBookRequest struct {
	DescribeBookRequest

	CreateBookRequest
}

func NewCreateBookRequest() *CreateBookRequest {
	return (&CreateBookRequest{}).SetIsSale(false)
}

func (r *CreateBookRequest) SetIsSale(v bool) *CreateBookRequest { // 链式调用
	r.IsSale = v
	return r
}
