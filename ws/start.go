/* ----------------------------------
*  @author suyame 2022-11-11 10:51:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"os"
)

func main() {
	r := InitRouter()
	r.Run(os.Args[1])
}
