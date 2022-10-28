/* ----------------------------------
*  @author suyame 2022-10-27 21:31:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package svg

import (
	"fmt"
	"strings"
)

func GenPath(userId, pageId int64) string {
	return fmt.Sprintf("svg/%d_%d.svg",
		userId, pageId)
}

func ParseFileName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
