/* ----------------------------------
*  @author suyame 2022-10-27 19:19:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package internal

import "sync"

type Value interface {
	Name() string
}

type NameSpace struct {
	sync.RWMutex
	name  string
	items map[string]Value
}

func New(name string) *NameSpace {
	return &NameSpace{
		name:  name,
		items: make(map[string]Value, 0),
	}
}

func (ns *NameSpace) Add(v Value) error {
	ns.Lock()
	defer ns.Unlock()
	ns.items[v.Name()] = v
	return nil
}
