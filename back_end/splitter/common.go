/* ----------------------------------
*  @author suyame 2022-11-11 11:11:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package splitter

import (
	"qiniu/config"
	"qiniu/splitter/loadbalance"
	"qiniu/splitter/server"
)

var (
	defaultWsServerMap = NewWsServerMap()
	servers            = []loadbalance.Server{}
	// 负载均衡(轮询)
	rr *loadbalance.RR
)

func init() {
	// 初始化servers
	for _, wsServer := range config.C.Servers {
		s := server.NewServer(wsServer.Host)
		// TODO ping 一下？
		servers = append(servers, s)
	}
	rr, _ = loadbalance.NewRR(servers)
}

// Allocate 分配ws服务器给指定页面
func Allocate(authorName string) (string, error) {
	key := authorName
	if host, ok := defaultWsServerMap.GetServerHost(key); ok {
		return host, nil
	}
	// 分配一个服务器给它
	s, err := rr.Do()
	if err != nil {
		return "", err
	}
	ws := s.(*server.Server)
	ws.Add()
	defaultWsServerMap.Put(key, ws)
	return ws.GetHost(), nil
}
