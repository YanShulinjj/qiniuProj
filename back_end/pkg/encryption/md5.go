/* ----------------------------------
*  @author suyame 2022-11-08 20:46:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package encryption

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5ByString(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr), nil
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
