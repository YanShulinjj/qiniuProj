/* ----------------------------------
*  @author suyame 2022-10-27 21:17:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package service

import (
	"sync"
	"ws/dao"
)

type userService struct{}

var (
	userServiceInstance *userService
	userServiceOnce     sync.Once
)

// NewUser 单例模式创建 UserService
func NewUser() *userService {
	userServiceOnce.Do(func() {
		userServiceInstance = &userService{}
	})

	return userServiceInstance
}

func (*userService) Add(username string) (int64, error) {
	return dao.NewUser().Create(username)
}

func (*userService) Drop(username string) error {
	return dao.NewUser().Drop(username)
}
