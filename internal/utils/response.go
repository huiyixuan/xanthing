package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通用响应结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 常用返回码常量
const (
	CodeSuccess            = 200
	CodeBadRequest         = 400
	CodeUnauthorized       = 401
	CodeForbidden          = 403
	CodeNotFound           = 404
	CodeInternalError      = 500
	CodeServiceUnavailable = 503
)

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  "success",
		Data: data,
	})
}

// 带自定义消息的成功响应
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

// 失败响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: map[string]interface{}{},
	})
}

// 错误请求响应
func BadRequest(c *gin.Context, msg string) {
	Fail(c, CodeBadRequest, msg)
}

// 未授权响应
func Unauthorized(c *gin.Context, msg string) {
	Fail(c, CodeUnauthorized, msg)
}

// 禁止访问响应
func Forbidden(c *gin.Context, msg string) {
	Fail(c, CodeForbidden, msg)
}

// 资源不存在响应
func NotFound(c *gin.Context, msg string) {
	Fail(c, CodeNotFound, msg)
}

// 内部错误响应
func InternalError(c *gin.Context, msg string) {
	Fail(c, CodeInternalError, msg)
}

// 服务不可用响应
func ServiceUnavailable(c *gin.Context, msg string) {
	Fail(c, CodeServiceUnavailable, msg)
}
