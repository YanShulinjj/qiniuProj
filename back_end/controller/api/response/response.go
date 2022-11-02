/* ----------------------------------
*  @author suyame 2022-10-27 22:14:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package response

import "qiniu/pkg/xerr"

type Status struct {
	Code    xerr.ErrCodeType `json:"status_code"`
	Message string           `json:"status_msg"`
}

type Register struct {
	Status
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type AddPage struct {
	Status
	PageIdx int64 `json:"page_idx"`
}

type Page struct {
	PageName string `json:"page_name"`
	SvgPath  string `json:"svg_path"`
}
type PageList struct {
	Status
	Items []*Page `json:"list"`
}
