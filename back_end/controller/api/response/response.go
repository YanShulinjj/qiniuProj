/* ----------------------------------
*  @author suyame 2022-10-27 22:14:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package response

import "ws/pkg/xerr"

type Status struct {
	Code    xerr.ErrCodeType `json:"status_code"`
	Message string           `json:"status_msg"`
}

type Register struct {
	Status
	UserID int64 `json:"user_id"`
}

type AddPage struct {
	Status
	PageIdx int64 `json:"page_idx"`
}
