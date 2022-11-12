/* ----------------------------------
*  @author suyame 2022-08-29 9:45:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package loadbalance

type Server interface {
	IsAlive() bool
}
