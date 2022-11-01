/* ----------------------------------
*  @author suyame 2022-10-26 20:32:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"net/http"
	"qiniu/ws/internal"

	"github.com/gin-gonic/gin"
)

func main() {

	// 设置调试模式或者发布模式必须是第一步！
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// // 注册中间件
	// r.Use(MiddleWare()) // 这个中间件注册在后面就无法起作用了，必须在前面调用。

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/wedraw", internal.WSHandler) // 每一个访问都会调用该路由对应的方法
	}

	bindAddress := ":9999"
	r.Run(bindAddress)
}

// func MiddleWare() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		fmt.Println("调用中间件，请求访问路径为：", ctx.Request.RequestURI)
// 	}
// }
