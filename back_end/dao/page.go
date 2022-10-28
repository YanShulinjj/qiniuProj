/* ----------------------------------
*  @author suyame 2022-10-27 20:55:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package dao

import (
	"github.com/pkg/errors"
	"sync"
	"ws/model"
	xerr2 "ws/pkg/xerr"
)

type pageDAO struct{}

var (
	pageDAOInstance *pageDAO
	pageDAOOnce     sync.Once
)

// NewUser 单例模式创建 userLoginDAO 实例
func NewPage() *pageDAO {
	pageDAOOnce.Do(func() {
		pageDAOInstance = &pageDAO{}
	})

	return pageDAOInstance
}

// Create 添加页面记录
func (u *pageDAO) Create(userId, pageIdx int64, svgPath string) (int64, error) {
	// 先判断该记录是否已经加入
	page, err := svcCtx.PageModel.FindOneByUserIdPageIdx(ctx, userId, pageIdx)
	if err == nil {
		return 0, errors.Wrapf(xerr2.NewErrCode(xerr2.PageExistedErr),
			"page_id:%s,err:%v", page.PageId, err)
	} else if err != model.ErrNotFound {
		// 出现未知错误
		return 0, err
	} else {
		// 新建
		page := model.Page{
			UserId:  userId,
			PageIdx: pageIdx,
			SvgPath: svgPath,
		}
		res, err := svcCtx.PageModel.Insert(ctx, &page)
		if err != nil {
			return 0, err
		}
		pageId, _ := res.LastInsertId()
		return pageId, nil
	}
}

// QueryPage 查询指定用户以及指定pageIdx的 矢量图路径
func (u *pageDAO) QueryPage(userId, pageIdx int64) (string, error) {
	// 查询page是否存在
	page, err := svcCtx.PageModel.FindOneByUserIdPageIdx(ctx, userId, pageIdx)
	if err != nil {
		if err == model.ErrNotFound {
			return "", errors.Wrapf(xerr2.NewErrCode(xerr2.PageNotExistErr),
				"page 不存在，err: %v", err)
		}
		return "", err
	}

	return page.SvgPath, nil
}

// UpdatePage 更新指定用户以及指定pageIdx的 矢量图路径
func (u *pageDAO) UpdatePage(userId, pageIdx int64, svgPath string) error {
	// 查询page是否存在
	page, err := svcCtx.PageModel.FindOneByUserIdPageIdx(ctx, userId, pageIdx)
	if err != nil {
		if err == model.ErrNotFound {
			return errors.Wrapf(xerr2.NewErrCode(xerr2.PageNotExistErr),
				"page 不存在，err: %v", err)
		}
		return err
	}

	page.SvgPath = svgPath

	err = svcCtx.PageModel.Update(ctx, page)

	return err
}

// Drop 删除指定用户以及指定pageIdx的矢量图记录
func (u *pageDAO) Drop(userId, pageIdx int64) (string, error) {
	// 查询page是否存在
	page, err := svcCtx.PageModel.FindOneByUserIdPageIdx(ctx, userId, pageIdx)
	if err != nil {
		if err == model.ErrNotFound {
			return "", errors.Wrapf(xerr2.NewErrCode(xerr2.PageNotExistErr),
				"page 不存在，err: %v", err)
		}
		return "", err
	}
	err = svcCtx.PageModel.Delete(ctx, page.PageId)
	return page.SvgPath, err
}
