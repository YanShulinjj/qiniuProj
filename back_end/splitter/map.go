/* ----------------------------------
*  @author suyame 2022-11-11 11:20:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package splitter

import (
	"qiniu/splitter/server"
	"sync"
)

// WsServerMap 并发安全映射
type WsServerMap struct {
	sync.RWMutex
	items map[string]*server.Server
}

func NewWsServerMap() *WsServerMap {
	return &WsServerMap{
		items: make(map[string]*server.Server),
	}
}

func (w *WsServerMap) Put(key string, server *server.Server) {
	w.Lock()
	defer w.Unlock()
	w.items[key] = server
}

func (w *WsServerMap) Get(key string) (*server.Server, bool) {
	w.RLock()
	defer w.RUnlock()
	server, ok := w.items[key]
	return server, ok
}

func (w *WsServerMap) GetServerHost(key string) (string, bool) {
	w.RLock()
	defer w.RUnlock()
	server, ok := w.items[key]
	if !ok {
		return "", ok
	}
	return server.GetHost(), ok
}
