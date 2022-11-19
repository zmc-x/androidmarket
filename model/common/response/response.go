package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ResultCode int         `json:"resultCode"` // 状态码
	Msg        string      `json:"msg"`        // 信息
	Data       interface{} `json:"data"`       // 数据部分，通过空接口实现对其他struct的使用
}

const (
	Error        = http.StatusInternalServerError // 报错
	Success      = http.StatusOK                  // 成功
	InvalidParam = http.StatusBadRequest          // 入参错误
	Unlogin      = http.StatusUnauthorized        // 未登录
)

// Result 接口的信息的返回
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ResultCode: code,
		Msg:        msg,
		Data:       data,
	})
}

// OkWithDetail 操作成功详细数据
func OkWithDetail(c *gin.Context, data interface{}, msg string) {
	Result(c, Success, msg, data)
}

// FailWithDetail 操作失败详细信息
func FailWithDetail(c *gin.Context, data interface{}, msg string) {
	Result(c, Error, msg, data)
}

// FailLogin 未登录
func FailLogin(c *gin.Context, msg string) {
	Result(c, Unlogin, msg, map[string]interface{}{})
}

// FailParam 参数错误
func FailParam(c *gin.Context) {
	Result(c, InvalidParam, "输入的参数出现错误", map[string]interface{}{})
}
