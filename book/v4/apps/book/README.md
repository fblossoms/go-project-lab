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

1. 开发业务功能
```go
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

```

2. 注册路由
```go
type BookApiHandler struct {
	ioc.ObjectImpl

	svc book.Service
}

// Name 就是 API 的资源名称
// api/book/v1/books
func (h *BookApiHandler) Name() string {
	return "books"
}

// Init 就是对象的初始化
// 主要初始化对象的属性
// 构造函数
func (h *BookApiHandler) Init() error {
	h.svc = book.GetService()

	// 本地依赖
	//r := server.Gin

	// 框架托管，通过容器获取Server对象
	// 获取的Gin Engine对象
	// RootRouter() 可能存在URL容器冲突的问题
	// ObjectRouter() 可以自动加上业务板块前缀或者对象名称避免冲突，格式为/<prefix>/<service_name>/<object_version>/<object_name>
	r := ioc_gin.ObjectRouter(h)
	r.GET("", h.queryBook)

	r.POST("", h.createBook)

	return nil
}

func init() {
	ioc.Api().Registry(&BookApiHandler{})
}
```

## 业务注册

每写完一个业务，就需要注册到注册表ioc（注册表）
```go

import (
	// api impl
	_ "go18/book/v4/apps/book/api"

	// service impl
	_ "go18/book/v4/apps/book/impl"
	_ "go18/book/v4/apps/comment/impl"
)
```

## 启动服务
```go
func main() {
	// ioc框架：加载对象，配置对象，注入对象
	//ioc.DevelopmentSetupWithPath("")

	//server.Gin.Run()
	//application.Get().AppName
	//http.Get().Host
	ioc.DevelopmentSetupWithPath("book/v4/application.toml")
	server.Run(context.Background())

	// ioc直接提供server，直接run
	// mcube包含一个gin. engine
	//server.Run(context.Background())
	//cmd.Start()
}
```

```sh
2026-05-08T11:21:58+08:00 WARN   config/vault/vault.go:113 > vault address is empty, skipping initialization app:simple_api group:default hostname:Fblossoms logger:vault
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

2026-05-08T11:21:58+08:00 INFO   config/gin/framework.go:41 > enable gin recovery app:simple_api group:default hostname:Fblossoms logger:gin_webframework
2026-05-08T11:21:58+08:00 WARN   config/datasource/grom.go:285 > password is empty for static credential mode app:simple_api group:default hostname:Fblossoms logger:datasource
2026-05-08T11:21:58+08:00 INFO   config/datasource/grom.go:287 > using static credentials from config file app:simple_api group:default hostname:Fblossoms logger:datasource
[GIN-debug] GET    /api/simple_api/1.0.0/books --> go18/book/v4/apps/book/api.(*BookApiHandler).queryBook-fm (4 handlers)
[GIN-debug] POST   /api/simple_api/1.0.0/books --> go18/book/v4/apps/book/api.(*BookApiHandler).createBook-fm (4 handlers)
2026-05-08T11:21:58+08:00 INFO   config/jsonrpc/service.go:114 > no reigstry service app:simple_api group:default hostname:Fblossoms logger:jsonrpc
2026-05-08T11:21:58+08:00 INFO   ioc/server/server.go:79 > loaded configs: [app.1.0.0 trace.1.0.0 log.1.0.0 vault.1.0.0 validator.1.0.0 gin_webframework.1.0.0 datasource.1.0.0 grpc.1.0.0 http.1.0.0] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:21:58+08:00 INFO   ioc/server/server.go:79 > loaded default: [] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:21:58+08:00 INFO   ioc/server/server.go:79 > loaded controllers: [book.1.0.0] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:21:58+08:00 INFO   ioc/server/server.go:79 > loaded apis: [books.1.0.0 jsonrpc.1.0.0] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:21:58+08:00 INFO   config/http/http.go:145 > HTTP服务启动成功, 监听地址: 127.0.0.1:8010 app:simple_api group:default hostname:Fblossoms logger:http
```

## 总结

业务分区框架，我们专注于业务对象的开发，mcube相当于一个工具箱，承接其他非业务的公共功能

## 其他非功能需求，开箱即用，比如健康检查health check，比如metrics

```go
	// 健康检查
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/gin"

	// 非业务模块
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/gin"
```

[metric.v1 books.v1 health.v1] metric, health 使用注入的对象
```sh
(base) PS C:\Users\flyfl\Desktop\go_18\book\v4> go run main.go -f C:\Users\flyfl\Desktop\go_18\book\v4\application.toml
2026-05-08T11:40:50+08:00 WARN   config/vault/vault.go:113 > vault address is empty, skipping initialization app:simple_api group:default hostname:Fblossoms logger:vault
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2026-05-08T11:40:50+08:00 INFO   config/gin/framework.go:41 > enable gin recovery app:simple_api group:default hostname:Fblossoms logger:gin_webframework
2026-05-08T11:40:50+08:00 WARN   config/datasource/grom.go:285 > password is empty for static credential mode app:simple_api group:default hostname:Fblossoms logger:datasource
2026-05-08T11:40:50+08:00 INFO   config/datasource/grom.go:287 > using static credentials from config file app:simple_api group:default hostname:Fblossoms logger:datasource
[GIN-debug] GET    /metrics/                 --> github.com/infraboard/mcube/v2/ioc/apps/metric/gin.(*ginHandler).Registry.func1 (5 handlers)
2026-05-08T11:40:50+08:00 INFO   metric/gin/metric.go:89 > Get the Metric using http://127.0.0.1:8010/metrics app:simple_api group:default hostname:Fblossoms logger:metric
[GIN-debug] GET    /healthz/                 --> github.com/infraboard/mcube/v2/ioc/apps/health/gin.(*HealthChecker).HealthHandleFunc-fm (5 handlers)
2026-05-08T11:40:50+08:00 INFO   health/gin/check.go:55 > Get the Health using http://127.0.0.1:8010/healthz app:simple_api group:default hostname:Fblossoms logger:health_check
[GIN-debug] GET    /api/simple_api/1.0.0/books --> go18/book/v4/apps/book/api.(*BookApiHandler).queryBook-fm (5 handlers)
[GIN-debug] POST   /api/simple_api/1.0.0/books --> go18/book/v4/apps/book/api.(*BookApiHandler).createBook-fm (5 handlers)
2026-05-08T11:40:50+08:00 INFO   config/jsonrpc/service.go:114 > no reigstry service app:simple_api group:default hostname:Fblossoms logger:jsonrpc
2026-05-08T11:40:50+08:00 INFO   ioc/server/server.go:79 > loaded configs: [app.1.0.0 trace.1.0.0 log.1.0.0 vault.1.0.0 validator.1.0.0 gin_webframework.1.0.0 datasource.1.0.0 grpc.1.0.0 http.1.0.0] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:40:50+08:00 INFO   ioc/server/server.go:79 > loaded default: [] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:40:50+08:00 INFO   ioc/server/server.go:79 > loaded controllers: [book.1.0.0] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:40:50+08:00 INFO   ioc/server/server.go:79 > loaded apis: [metric.v1 health.1.0.0 books.1.0.0 jsonrpc.1.0.0] app:simple_api group:default hostname:Fblossoms logger:server
2026-05-08T11:40:50+08:00 INFO   config/http/http.go:145 > HTTP服务启动成功, 监听地址: 127.0.0.1:8010 app:simple_api group:default hostname:Fblossoms logger:http
```