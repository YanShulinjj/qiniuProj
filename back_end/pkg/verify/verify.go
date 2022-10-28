/* ----------------------------------
*  @author suyame 2022-10-27 21:59:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package verify

import "regexp"

func UserName(name string, ipMode bool) bool {
	if !ipMode {
		return len(name) != 0
	}

	// 如果是ip模式（暂时支持IPV4）
	reg, _ := regexp.Compile("^(?:(?:1[0-9][0-9]\\.)|(?:2[0-4][0-9]\\.)|(?:25[0-5]\\.)|(?:[1-9][0-9]\\.)|(?:[0-9]\\.)){3}(?:(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])|(?:[1-9][0-9])|(?:[0-9]))$")
	return reg.MatchString(name)
}
