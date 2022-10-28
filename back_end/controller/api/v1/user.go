/* ----------------------------------
*  @author suyame 2022-10-27 21:56:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws/controller/api/response"
	"ws/pkg/xerr"
	"ws/service"
)

type userController struct{}

// NewUserController 是 userController 的构造器
// 返回一个 userController 的指针
func NewUserController() *userController {
	return &userController{}
}

func (*userController) Register(c *gin.Context) {
	username := c.Query("username")
	userId, err := service.NewUser().Add(username)
	if err != nil {
		c.JSON(http.StatusOK, response.Register{
			Status: response.Status{
				xerr.ReuqestParamErr,
				"注册用户出错",
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.Register{
		Status: response.Status{},
		UserID: userId,
	})
}
