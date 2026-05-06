package api

import (
	"go18/book/v4/apps/book"

	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"

	"github.com/infraboard/mcube/v2/ioc"
)

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
