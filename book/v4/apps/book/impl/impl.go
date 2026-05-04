package impl

import (
	"go18/book/v4/apps/book"

	"github.com/infraboard/mcube/v2/ioc"
)

// 如何知道有没有实现呢？通过类型约束
// var _ book.Service = &BookServiceImpl{}
// &BookServiceImpl 的 nil（空指针）的对象，减少内存开销
var _ book.Service = (*BookServiceImpl)(nil)

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
