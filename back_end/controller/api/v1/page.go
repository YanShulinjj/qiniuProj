/* ----------------------------------
*  @author suyame 2022-10-27 22:23:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	pagename := c.DefaultQuery("pagename", "unnamed")
	svgPath, err := service.NewPage().Add(username, pagename)
	if err != nil {
		c.JSON(http.StatusOK, response.AddPage{
			Status: response.Status{
				xerr.ReuqestParamErr,
				"添加页面出错！" + err.Error(),
			},
		})
		return
	}
	// // svgPath
	// svgPath := filepath.Join(config.C.SVGPATH, svgFileName)
	// // 保存到服务器
	// fmt.Println("Saving ....", svgPath)

	c.JSON(http.StatusOK, response.AddPage{
		Status:   response.Status{},
		PageName: pagename,
		SvgPath:  svgPath,
	})
}

func (*pageController) PageList(c *gin.Context) {
	username := c.Query("author")

	pages, err := service.NewPage().QueryMany(username)
	if err != nil {
		c.JSON(http.StatusNotFound, response.PageList{
			Status: response.Status{
				xerr.ReuqestParamErr,
				"添加页面出错！",
			},
		})
		return
	}
	var items []*response.Page
	for _, page := range pages {
		items = append(items,
			&response.Page{
				PageName: page[0],
				SvgPath:  page[1],
			})
	}
	c.JSON(http.StatusOK, response.PageList{
		Items: items,
	})
}

func (*pageController) UploadSVG(c *gin.Context) {
	// 从客户端传输svg文件
	username := c.PostForm("username")
	filename := c.PostForm("pagename")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, response.Status{
			xerr.ReuqestParamErr,
			"上传文件错误，未识别出文件",
		})
		return
	}
	// 保存到服务端
	saveDir := filepath.Join(config.C.SVGPATH, username)

	// 如果文件夹不存在
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		err = os.Mkdir(saveDir, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, response.Status{
				xerr.ReuqestParamErr,
				"服务端创建文件夹失败",
			})
			return
		}
	}
	saveFile := filepath.Join(saveDir, filename)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, response.Status{
			xerr.ReuqestParamErr,
			"上传文件错误，未识别出文件" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Status{
		0,
		"ok",
	})

}

// GetPage 根据用户和pagename 获取svgpath
func (*pageController) GetPage(c *gin.Context) {
	username := c.Query("username")
	pagename := c.Query("pagename")

	path, err := service.NewPage().QueryOne(username, pagename)
	if err != nil {
		c.JSON(http.StatusOK, response.Status{
			xerr.ReuqestParamErr,
			"获取page错误，" + err.Error(),
		})
		return
	}
	//
	c.JSON(http.StatusOK, response.PageInfo{
		Page: response.Page{
			PageName: pagename,
			SvgPath:  "http://" + path,
		},
	})
}
