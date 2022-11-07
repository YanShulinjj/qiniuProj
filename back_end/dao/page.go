/* ----------------------------------
*  @author suyame 2022-10-27 20:55:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package dao

import (
	"github.com/pkg/errors"
	"qiniu/model"
	xerr2 "qiniu/pkg/xerr"
	"sync"
)

type pageDAO struct{}

var (
	pageDAOInstance *pageDAO
	pageDAOOnce     sync.Once
)

// NewPage 单例模式创建 pageDAO 实例
func NewPage() *pageDAO {
	pageDAOOnce.Do(func() {
		pageDAOInstance = &pageDAO{}
	})

	return pageDAOInstance
}

// Create 添加页面记录
func (p *pageDAO) Create(userId, pageIdx int64, pageName, svgPath string) (int64, error) {
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
			UserId:   userId,
			PageIdx:  pageIdx,
			PageName: pageName,
			SvgPath:  svgPath,
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
func (p *pageDAO) QueryPage(userId, pageIdx int64) (string, string, error) {
	// 查询page是否存在
	page, err := svcCtx.PageModel.FindOneByUserIdPageIdx(ctx, userId, pageIdx)
	if err != nil {
		if err == model.ErrNotFound {
			return "", "", errors.Wrapf(xerr2.NewErrCode(xerr2.PageNotExistErr),
				"page 不存在，err: %v", err)
		}
		return "", "", err
	}

	return page.PageName, page.SvgPath, nil
}

// QueryPage 查询指定用户以及指定pageIdx的 矢量图路径
func (p *pageDAO) QueryPages(userId int64) ([][]string, error) {
	// 查询page是否存在
	pages, err := svcCtx.PageModel.FindMany(ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr2.NewErrCode(xerr2.PageNotExistErr),
				"page 不存在，err: %v", err)
		}
		return nil, err
	}
	resp := [][]string{}
	for _, page := range pages {
		resp = append(resp, []string{page.PageName, page.SvgPath})
	}
	return resp, nil
}

// UpdatePage 更新指定用户以及指定pageIdx的 矢量图路径
func (p *pageDAO) UpdatePage(userId, pageIdx int64, pageName, svgPath string) error {
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
	page.PageName = pageName

	err = svcCtx.PageModel.Update(ctx, page)

	return err
}

// Drop 删除指定用户以及指定pageIdx的矢量图记录
func (p *pageDAO) Drop(userId, pageIdx int64) (string, error) {
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

// QueryPageByName 查询指定的userid和pageName 是否存在
func (p *pageDAO) QueryPageByName(userId int64, pageName string) (string, bool) {
	// 查询page是否存在
	page, err := svcCtx.PageModel.FindOneByUserIdPageName(ctx, userId, pageName)
	if err != nil {
		return "", false
	}
	return page.SvgPath, true
}
