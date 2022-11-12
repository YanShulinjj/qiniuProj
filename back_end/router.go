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
	"qiniu/splitter"
	"strings"
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
		rwMode := c.DefaultQuery("rw", "0")
		Ip := c.RemoteIP()
		userName, _ := encryption.Md5ByString(Ip)
		authorName := c.Query("author")
		if len(authorName) == 0 {
			// 新建一个用户
			service.NewUser().Add(userName)
			authorName = userName
		}
		// 如果页面还未添加
		service.NewPage().Add(authorName, pageName)

		// 分配一个ws服务器给当前页面
		wsHost, err := splitter.Allocate(authorName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "faild",
				"msg":    "获取ws服务器失败," + err.Error(),
			})
		}
		// 因为单机测试使用 多个端口模拟多个ws服务器

		ip := strings.Split(c.Request.Host, ":")[0]
		port := strings.Split(wsHost, ":")[1]

		c.HTML(http.StatusOK, "index.html", gin.H{
			"pageName":   pageName,
			"userName":   userName,
			"authorName": authorName,
			"hostAddr":   c.Request.Host,
			"wsAddr":     ip + ":" + port,
			"rwMode":     rwMode,
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

	return r
}
