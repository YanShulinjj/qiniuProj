/* ----------------------------------
*  @author suyame 2022-10-27 21:34:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package config

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type WsServerConfig struct {
	Host string
}
type Config struct {
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Host       string
	Port       string
	SVGPATH    string
	Servers    []WsServerConfig
}

var C Config

func init() {
	var configFile = flag.String("f", "etc/config.yaml", "the config file")
	conf.MustLoad(*configFile, &C)
}
