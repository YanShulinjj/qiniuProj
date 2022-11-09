/* ----------------------------------
*  @author suyame 2022-11-01 21:45:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "qiniu/controller/api/v1"
	"qiniu/pkg/encryption"
	"qiniu/service"
	"qiniu/ws"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})

	// 将指定目录挂在web
	r.StaticFS("/public", http.Dir("../front_end/public"))
	r.StaticFS("/data", http.Dir("./data"))
	r.LoadHTMLFiles("../front_end/index.html")
	r.GET("/qiniu", func(c *gin.Context) {
		pageName := c.DefaultQuery("page", "1")
		Ip := c.RemoteIP()
		userName, _ := encryption.Md5ByString(Ip)
		authorName := c.Query("author")
		if len(authorName) == 0 {
			// 新建一个用户
			service.NewUser().Add(userName)
			authorName = userName
		}
		// 如果页面还未添加
		service.NewPage().Add(userName, pageName)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"pageName":   pageName,
			"userName":   userName,
			"authorName": authorName,
			"hostAddr":   c.Request.Host,
		})
	})
	// TODO
	backend := r.Group("/backend/")
	{
		backend.GET("/page/get", v1.PageController.GetPage)
		backend.GET("/page/add", v1.PageController.Add)
		backend.GET("/page/list", v1.PageController.PageList)
		backend.POST("/page/upload", v1.PageController.UploadSVG)
		backend.GET("/user/add", v1.UserController.Register)

	}

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/wedraw", ws.WSHandler) // 每一个访问都会调用该路由对应的方法
		wsGroup.GET("/statue", ws.SyncHandler)
	}

	return r
}
