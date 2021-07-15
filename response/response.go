package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jormin/gobox/errs"
	"github.com/jormin/golog/log"
)

// HTTP响应
type APIResponse struct {
	Code    int         `json:"code" desc:"业务状态码"`         // 状态码
	Message string      `json:"message" desc:"业务状态信息"`     // 状态信息
	Data    interface{} `json:"data,omitempty"desc:"业务数据"` // 状态信息
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Data(status, "application/json; charset=utf-8", buf)
}

// 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	response := APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    v,
	}
	ResJSON(c, http.StatusOK, response)
}

// 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	response := APIResponse{
		Code:    errs.ErrUnknown.Code,
		Message: errs.ErrUnknown.Error(),
	}
	if err != nil {
		response.Message = err.Error()
		if e, ok := err.(*errs.Error); ok {
			response.Code = e.Code
		}
	}
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	if statusCode != http.StatusOK {
		log.Warn("res error: %s", err.Error())
	}
	ResJSON(c, statusCode, response)
}
