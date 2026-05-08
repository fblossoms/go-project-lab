package main

import (
	"context"
	_ "go18/book/v4/apps/book/api"
	_ "go18/book/v4/apps/book/impl"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/server"

	// 健康检查
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/gin"

	// 非业务模块
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/gin"
)

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
