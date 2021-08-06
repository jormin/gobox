package errs

import (
	"encoding/json"
	"fmt"
)

var (
	ErrUnknown = NewError(99999, "未知错误")
	ErrParams  = NewError(10000, "参数错误")
)

// Error 定义业务错误
type Error struct {
	Code    int    `json:"code"`              // 错误码
	Message string `json:"message,omitempty"` // 错误信息
}

// 错误输出
func (r *Error) Error() string {
	b, _ := json.Marshal(r)
	return fmt.Sprintf("%s", b)
}

// 新建错误
func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
