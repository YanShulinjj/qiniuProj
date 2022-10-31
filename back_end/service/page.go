/* ----------------------------------
*  @author suyame 2022-10-27 21:17:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package service

import (
	"qiniu/dao"
	"qiniu/pkg/svg"
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
func (*pageService) Add(username string) (string, int64, error) {
	// 首先根据username获取userid和pageIdx
	userid, pageIdx, err := dao.NewUser().Query(username)
	if err != nil {
		return "", 0, err
	}
	pageIdx += 1
	// 生成一个svgPath路径
	svgPath := svg.GenPath(userid, pageIdx)
	_, err = dao.NewPage().Create(userid, pageIdx, svgPath)
	if err != nil {
		return "", 0, err
	}
	// 更新pageNum
	err = dao.NewUser().UpdatePageNum(username, pageIdx)
	return svgPath, pageIdx, nil
}

// Drop 指定用户新建一个page
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
