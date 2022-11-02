/* ----------------------------------
*  @author suyame 2022-11-01 21:09:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package ws

import "sync"

type ManagerGroup struct {
	sync.RWMutex
	items map[string]*Manager
}

var defaultMangerGroup = NewMangerGroup()

func NewMangerGroup() *ManagerGroup {
	return &ManagerGroup{
		items: make(map[string]*Manager, 0),
	}
}

func (mg *ManagerGroup) Put(pageName string, m *Manager) {
	mg.Lock()
	defer mg.Unlock()
	mg.items[pageName] = m
}

func (mg *ManagerGroup) Get(pageName string) (*Manager, bool) {
	mg.RLock()
	defer mg.RUnlock()
	v, ok := mg.items[pageName]
	return v, ok
}

func (mg *ManagerGroup) IsExist(pageName string) bool {
	mg.RLock()
	defer mg.Unlock()
	_, ok := mg.items[pageName]
	return ok
}
