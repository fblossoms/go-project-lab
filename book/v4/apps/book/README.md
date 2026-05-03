# Book业务分区

定义Book业务逻辑

业务功能：CRUD
1. 创建（录入）书籍
2. Book列表查询
3. Book详情查询
4. Book更新
5. Book删除

通过Go语言里面的接口 来定义描述业务功能

```go
// Service Book的业务定义
type Service interface {
	// CreateBook 1. 创建（录入）书籍
	CreateBook(ctx context.Context, request *CreateBookRequest) (Book, error)

	// QueryBook 2. Book列表查询
	QueryBook(ctx context.Context, request *QueryBookRequest) (*types.Set[*Book], error)

	// DescribeBook 3. Book详情查询
	DescribeBook(ctx context.Context, request *DescribeBookRequest) (*Book, error)

	// UpdateBook 4. Book更新
	UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error)

	// DeleteBook  5. Book删除
	DeleteBook(ctx context.Context, request *DeleteBookRequest) (*Book, error)
}
```

## 业务的具体实现（使用TDD思想：Test Drive Develop测试驱动开发）

1. BookServiceImpl

    ```go
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
    ```

2. 编写单元测试