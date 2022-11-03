/* ----------------------------------
*  @author suyame 2022-11-01 20:42:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import "qiniu/config"

func main() {
	r := InitRouter()
	r.Run(config.C.Host + config.C.Port)
}
