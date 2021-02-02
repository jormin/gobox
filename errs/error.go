package errs

import (
	"errors"
	"fmt"
)

var (
	ErrUnknown = NewErrorMsg(99999, "未知错误")
)

// Error 定义业务错误
type Error struct {
	Code int   `json:"code"`          // 错误码
	Err  error `json:"err,omitempty"` // 响应错误
}

// 错误输出
func (r *Error) Error() string {
	if r.Err != nil {
		return fmt.Sprintf("[%d]%s", r.Code, r.Err.Error())
	}
	return fmt.Sprintf("[%d]未定义错误码", r.Code)
}

// 新建错误
func NewError(code int, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

// 新建错误
func NewErrorMsg(code int, msg string) *Error {
	return &Error{
		Code: code,
		Err:  errors.New(msg),
	}
}
