package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 添加监听路由
	server := gin.Default() // 可以理解为这是对HTTP Server的一个包装，让其提供更完善的HTTP协议与服务

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Gin Hello, world!")
	})

	err := server.Run(":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
