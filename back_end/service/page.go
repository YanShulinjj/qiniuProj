/* ----------------------------------
*  @author suyame 2022-10-27 21:17:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package service

import (
	"github.com/pkg/errors"
	"qiniu/config"
	"qiniu/dao"
	"qiniu/pkg/svg"
	xerr2 "qiniu/pkg/xerr"
	"sync"
)

type pageService struct{}

var (
	pageServiceInstance *pageService
	pageServiceOnce     sync.Once
)

// NewPage 单例模式创建 UserService
func NewPage() *pageService {
	pageServiceOnce.Do(func() {
		pageServiceInstance = &pageService{}
	})

	return pageServiceInstance
}

// Add 指定用户新建一个page
func (*pageService) Add(username, pagename string) (string, error) {
	// 首先根据username获取userid和pageIdx
	userid, pageIdx, err := dao.NewUser().Query(username)
	if err != nil {
		return "", err
	}
	// 为防止用户创建同名
	if _, ok := dao.NewPage().QueryPageByName(userid, pagename); ok {
		return "", errors.Wrapf(xerr2.NewErrCode(xerr2.PageExistedErr),
			"page 已经存在, 不允许再次创建")
	}
	pageIdx += 1
	// 生成一个svgPath路径
	svgPath := svg.GenPath(config.C.Host, config.C.Port, username, pagename)
	_, err = dao.NewPage().Create(userid, pageIdx, pagename, svgPath)
	if err != nil {
		return "", err
	}
	// 更新pageNum
	err = dao.NewUser().UpdatePageNum(username, pageIdx)
	return svgPath, nil
}

// Drop 删除指定用户的一个page
func (*pageService) Drop(username string, pageIdx int64) (string, error) {
	// 首先根据username获取userid
	userid, _, err := dao.NewUser().Query(username)
	if err != nil {
		return "", err
	}
	avgPath, err := dao.NewPage().Drop(userid, pageIdx)
	if err != nil {
		return "", err
	}
	filename := svg.ParseFileName(avgPath)
	return filename, err
}

func (*pageService) QueryMany(username string) ([][]string, error) {
	// 首先根据username获取userid
	userid, _, err := dao.NewUser().Query(username)
	if err != nil {
		return nil, err
	}

	pages, err := dao.NewPage().QueryPages(userid)
	return pages, err
}

func (*pageService) QueryOne(username string, pagename string) (string, error) {
	// 首先根据username获取userid
	userid, _, err := dao.NewUser().Query(username)
	if err != nil {
		return "", err
	}
	SvgPath, ok := dao.NewPage().QueryPageByName(userid, pagename)
	if !ok {
		return "", errors.Wrapf(xerr2.NewErrCode(xerr2.PageNotExistErr),
			"page 不存在")
	}
	return SvgPath, nil
}
