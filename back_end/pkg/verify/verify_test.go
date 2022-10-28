/* ----------------------------------
*  @author suyame 2022-10-27 22:06:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package verify

import "testing"

func TestUserName(t *testing.T) {
	cases := []string{
		"1.1.1.1",
		"10.10.10.10",
		"255.255.255.255",
		"0.0.0.0",
		"256.0.0.0.1",
		"-1.1.1.1",
		"2222.0.0.1",
		"0.0.0.0.0",
	}
	expects := []bool{
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
	}

	for i, c := range cases {
		if get := UserName(c, true); get != expects[i] {
			t.Errorf("IPV4 format test err, expect: %v, get: %v\n", expects[i], get)
		}
	}
}
