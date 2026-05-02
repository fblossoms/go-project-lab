package exception

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Recovery 自定义异常处理
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		fmt.Println(err)
		c.JSON(500, "服务器异常")
		c.Abort()
	})
}
