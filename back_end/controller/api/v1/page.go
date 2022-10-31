/* ----------------------------------
*  @author suyame 2022-10-27 22:23:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"qiniu/config"
	"qiniu/controller/api/response"
	"qiniu/pkg/xerr"
	"qiniu/service"
)

type pageController struct{}

// NewPageController 是 pageController 的构造器
// 返回一个 pageController 的指针
func NewPageController() *pageController {
	return &pageController{}
}

func (*pageController) Add(c *gin.Context) {
	username := c.Query("username")
	svgFileName, pageIdx, err := service.NewPage().Add(username)
	if err != nil {
		c.JSON(http.StatusOK, response.AddPage{
			Status: response.Status{
				xerr.ReuqestParamErr,
				"添加页面出错！",
			},
		})
		return
	}

	// svgPath
	svgPath := filepath.Join(config.C.SVGPATH, svgFileName)
	// 保存到服务器
	fmt.Println("Saving ....", svgPath)

	c.JSON(http.StatusOK, response.AddPage{
		Status:  response.Status{},
		PageIdx: pageIdx,
	})
}
