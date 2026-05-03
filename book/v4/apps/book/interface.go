package book

import "context"

// Service Book的业务定义
type Service interface {
	// CreateBook 1. 创建（录入）书籍
	CreateBook(ctx context.Context, request *CreateBookRequest) (Book, error)

	// 2. Book列表查询
	QueryBook(ctx context.Context, request *QueryBookRequest) BookSet
	// 3. Book详情查询
	// 4. Book更新
	// 5. Book删除
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

type QueryBookRequest struct {
}
