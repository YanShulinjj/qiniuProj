/* ----------------------------------
*  @author suyame 2022-11-11 10:56:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package conf

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type ServerConfig struct {
	Host string
}

type Config struct {
	Servers []ServerConfig
}

var C Config

func init() {
	var configFile = flag.String("f", "etc/config.yaml", "the config file")
	conf.MustLoad(*configFile, &C)
}
