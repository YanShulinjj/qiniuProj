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

func GenPath(host, port, username, pagename string) string {
	return fmt.Sprintf("%s%s/data/svg/%s/%s",
		host, port, username, pagename)
}

func ParseFileName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
