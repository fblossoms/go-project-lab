package response

import (
	"go18/book/v3/exception"

	"github.com/gin-gonic/gin"
)

// OK 当请求成功时，我们应用返回的数据
// 1. {code: 0, data: {}}
// 2. 正常直接返回数据，Restful接口 怎么知道这些请求是成功还是失败？通过HttpCode判断，如2xx
// 如果后面所有的返回数据要经过特殊处理，都在这个函数内进行扩展，方便维护，比如 数据脱敏
func OK(ctx *gin.Context, data any) {
	ctx.JSON(200, data)
	ctx.Abort() // 中断
}

// Failed 当请求失败时，我们返回数据格式
// 1. {code: xxx（不为0）, data: null, message: "（错误信息）"
// 2. 通过HttpCode判断，若非2xx，就返回自定义的异常
func Failed(ctx *gin.Context, err error) {
	// 断言异常
	e, ok := err.(*exception.ApiException)
	// 如果是业务异常
	if ok {
		ctx.JSON(e.HttpCode, e)
		ctx.Abort()
		return
	}
	// 如果是非业务异常
	ctx.JSON(500, exception.NewApiException(500, err.Error()))
	ctx.Abort() // 中断

	//ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
	//ctx.Abort() // 中断
}
