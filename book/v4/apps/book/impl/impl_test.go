package impl_test

import (
	"context"
	"go18/book/v4/apps/book"
	"go18/book/v4/apps/book/test"
)

var ctx = context.Background()
var svc book.Service // 声明对象

func init() {
	// 引入包后自动执行的逻辑
	// 工具对象的初始化，需要绝对路径
	//ioc.DevelopmentSetupWithPath("C:/Users/flyfl/Desktop/go_18/book/v4/application.toml")

	test.DevelopmentSet()

	svc = book.GetService() // 获取对象
}
