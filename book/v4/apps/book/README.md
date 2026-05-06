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

   ```go
    func TestCreateBook(t *testing.T) {
    	req := book.NewCreateBookRequest()
    	req.SetIsSale(true)
    	req.Title = "Go语言V4.0"
    	req.Author = "will"
    	req.Price = 99.9
    	ins, err := svc.CreateBook(ctx, req)
    	if err != nil {
    		t.Fatal(err)
    	}
    	t.Log(ins)
    }
   ```

3. 业务对象注册（注册到ioc controller）

注册到全局变量池

如果是手动维护，不方便
   ```sh
pkg globle

bookController = xxx
commentController = xxx
...
   ```

通过容器来维护对象

   ```go
// BookServiceImpl Book业务的具体实现
// 作为具体实现就会有很多种情况
// 可以添加多种实现：例如项目后续更新可能使用到多种技术栈，就会产生多种情况
type BookServiceImpl struct {
	ioc.ObjectImpl // 注册对象
}

// Name 重写 Name 方法，返回控制器（对象）的名称
// 当前以服务名称作为实现名称
// 当前的 BookServiceImpl 是 Service book.APP_NAME 的一个具体实现
func (s *BookServiceImpl) Name() string {
	return book.APP_NAME
}

// 写好一个业务对象（业务实现），就把放入（注册）一个容器（公共空间，或全局变量）中
// mcube提供了该空间 ioc.Controller().Registry 把对象注册过去
// 需提供对象的名称、初始化函数、初始化方法
func init() {
	ioc.Controller().Registry(&BookServiceImpl{})
}
   ```
   
## 面向接口

把对象取出来，断言满足业务接口，然后我们以接口的方式来使用

```go
// APP_NAME 当前以服务名称作为实现名称
const (
APP_NAME = "book"
)

func GetService() Service {
return ioc.Controller().Get(APP_NAME).(Service) // 取出来并断言成接口
}
```

第三方的模块，可以依赖接口进行开发
```go
func (c CommentServiceImpl) AddComment(ctx context.Context, request *comment.AddCommentRequest) (*comment.Comment, error) {
// 问题：能不能直接依赖Book.Service的具体实现？
// 不能，强依赖具体业务实现，若后续技术栈升级，工程量激增
// 应依赖接口，面向接口编程，不依赖业务实现
//(&impl.BookServiceImpl{}).DescribeBook(ctx, nil)
book.GetService().DescribeBook(ctx)
}

```

## 开发API

接口是需求，对业务进行设计，可以选择把这些能力以哪种接口的形式对外提供服务

1. 不对外提供接口，仅作为其他业务的依赖
2. （API）对外提供HTTP接口，RESTful接口
3. （API）对内提供接口（JSON RPC/gRPC）