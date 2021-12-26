package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	// 当客户端有请求来之后, 先执行这个函数

	return func(c *gin.Context) {
		fmt.Println("MiddleWare: 中间件开始执行")
		c.Set("request", "中间件") // 在gin.Context中设置一个值(这个与中间件无关, 只是为了演示)

		// 执行到这个函数时, 会跳转到main函数中, 执行客户端对应的请求回调函数
		c.Next()
		// 执行完对应的回调函数之后, 继续回到这个地方进行执行(但是响应还没有返回给客户端)

		// 获取相应状态码, 并打印相关信息
		status := c.Writer.Status()
		fmt.Println("MiddleWare: 中间件执行结束, status: ", status)
	}

	// 当中间件执行完之后, 才真正把响应返回给客户端
}
