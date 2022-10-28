/* ----------------------------------
*  @author suyame 2022-10-27 20:13:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package dao

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ws/config"
	"ws/model"
)

type ServiceContext struct {
	Config config.Config
	// 添加user model 依赖
	UserModel model.UserModel
	// 添加page model 依赖
	PageModel model.PageModel
}

var svcCtx *ServiceContext
var ctx context.Context

func init() {
	svcCtx = newServiceContext(config.C)
}

func newServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
		PageModel: model.NewPageModel(sqlConn, c.CacheRedis),
	}
}
