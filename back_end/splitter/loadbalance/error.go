/* ----------------------------------
*  @author suyame 2022-08-29 9:48:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package loadbalance

import "errors"

var (
	NoAliveServerErr   = errors.New("No server can accpet this request!")
	WeightsNotMatchErr = errors.New("weights not valid!")
)
