/* ----------------------------------
*  @author suyame 2022-11-02 19:51:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package name

import (
	"fmt"
	"sync/atomic"
)

var seq int64

func GetDefaultName() string {
	atomic.AddInt64(&seq, 1)
	return fmt.Sprintf("unnamed_%d", seq)
}
