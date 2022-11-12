/* ----------------------------------
*  @author suyame 2022-11-11 10:52:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws/internal"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/wedraw", internal.WSHandler) // 每一个访问都会调用该路由对应的方法
		wsGroup.GET("/statue", internal.SyncHandler)
	}

	return r
}
