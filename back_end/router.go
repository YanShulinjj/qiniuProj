/* ----------------------------------
*  @author suyame 2022-11-01 21:45:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 将指定目录挂在web
	r.StaticFS("/public", http.Dir("../front_end/public"))

	r.LoadHTMLFiles("../front_end/index.html")
	r.GET("/qiniu", func(c *gin.Context) {
		pageName := c.Query("page")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"pageName": pageName,
		})
	})
	// // TODO
	// backend := r.Group("/backend/")
	// {
	// 	backend.GET("/page/add", v1.PageController.Add)
	// 	backend.POST("/user/add", v1.UserController.Register)
	// }
	return r
}
