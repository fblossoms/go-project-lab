package test

import (
	"github.com/infraboard/mcube/v2/ioc"

	// 把要注册的对象准备好，Book，Comment
	_ "go18/book/v4/apps/book/impl"
	_ "go18/book/v4/apps/comment/impl"
)

func DevelopmentSet() {
	ioc.DevelopmentSetupWithPath("C:/Users/flyfl/Desktop/go_18/book/v4/application.toml")
}
