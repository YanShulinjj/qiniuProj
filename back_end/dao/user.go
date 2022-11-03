/* ----------------------------------
*  @author suyame 2022-10-27 20:19:00
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

type userDAO struct{}

var (
	userDAOInstance *userDAO
	userDAOOnce     sync.Once
)

// NewUser 单例模式创建 userLoginDAO 实例
func NewUser() *userDAO {
	userDAOOnce.Do(func() {
		userDAOInstance = &userDAO{}
	})

	return userDAOInstance
}

// Create 添加新用户
func (u *userDAO) Create(username string) (int64, error) {
	// 先判断，如果用户名已经存在，返回
	user, err := svcCtx.UserModel.FindOneByUserIp(ctx, username)
	if err == nil {
		// return 0, errors.Wrapf(xerr2.NewErrCode(xerr2.UserExistedErr),
		// 	"username 已经存在啦，user_id:%s,err:%v", user.UserId, err)

		// 如果已经存在 直接返回userId
		return user.UserId, nil
	} else if err != model.ErrNotFound {
		// 出现未知错误
		return 0, err
	} else {
		// 新建
		user := model.User{
			UserIp: username,
		}
		res, err := svcCtx.UserModel.Insert(ctx, &user)
		if err != nil {
			return 0, err
		}
		userId, _ := res.LastInsertId()
		return userId, nil
	}
}

func (u *userDAO) Query(username string) (int64, int64, error) {
	// 查询用户是否存在
	user, err := svcCtx.UserModel.FindOneByUserIp(ctx, username)
	if err != nil {
		if err == model.ErrNotFound {
			return 0, 0, errors.Wrapf(xerr2.NewErrCode(xerr2.UserNotExistErr),
				"username 不存在，err: %v", err)
		}
		return 0, 0, err
	}
	return user.UserId, user.NextPage, nil
}

// QueryPageNum 查询下一个新建page索引
func (u *userDAO) QueryPageNum(username string) (int64, error) {
	// 查询用户是否存在
	user, err := svcCtx.UserModel.FindOneByUserIp(ctx, username)
	if err != nil {
		if err == model.ErrNotFound {
			return 0, errors.Wrapf(xerr2.NewErrCode(xerr2.UserNotExistErr),
				"username 不存在，err: %v", err)
		}
		return 0, err
	}
	return user.NextPage, nil
}

// UpdatePageNum 更新指定用户的pagenum字段
func (u *userDAO) UpdatePageNum(username string, pageNum int64) error {
	// 查询用户是否存在
	user, err := svcCtx.UserModel.FindOneByUserIp(ctx, username)
	if err != nil {
		if err == model.ErrNotFound {
			return errors.Wrapf(xerr2.NewErrCode(xerr2.UserNotExistErr),
				"username 不存在，err: %v", err)
		}
		return err
	}
	user.NextPage = pageNum
	err = svcCtx.UserModel.Update(ctx, user)
	return err
}

// Drop 删除指定用户
func (u *userDAO) Drop(username string) error {
	// 查询用户是否存在
	user, err := svcCtx.UserModel.FindOneByUserIp(ctx, username)
	if err != nil {
		if err == model.ErrNotFound {
			return errors.Wrapf(xerr2.NewErrCode(xerr2.UserNotExistErr),
				"username 不存在，err: %v", err)
		}
		return err
	}
	err = svcCtx.UserModel.Delete(ctx, user.UserId)
	return err
}
