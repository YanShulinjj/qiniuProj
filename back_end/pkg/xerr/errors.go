package xerr

import (
	"fmt"
)

type CodeError struct {
	errCode ErrCodeType
	errMsg  string
}

// errorCode
func (e *CodeError) GetErrCode() ErrCodeType {
	return e.errCode
}

// errMsg
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode ErrCodeType, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewErrCode(errCode ErrCodeType) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: ServerCommonErr, errMsg: errMsg}
}
