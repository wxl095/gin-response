package response

import (
	"github.com/gin-gonic/gin"
	"github.com/wxl095/gin-response/errcode"
	"net/http"
)

type ErrCodeInterface interface {
	String() string
}

type Response struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    any    `json:"data"`
}
type NoDataResponse struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

func NewNoDataResponse() *NoDataResponse {
	return &NoDataResponse{}
}

func NewResponse() *Response {
	return &Response{
		ErrCode: int(errcode.OK),
		ErrMsg:  ErrCodeInterface(errcode.OK).String(),
		Data:    nil,
	}
}

type Wrapper struct {
	ctx    *gin.Context
	render func(code int, data any)
}

// DefaultWrapContext 默认装饰使用json
func DefaultWrapContext(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx: ctx, render: ctx.JSON}
}

func WrapContextWithXml(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx: ctx, render: ctx.XML}
}

// Success func (w *Wrapper) Xml(response *Response) {
//	w.ctx.XML(http.StatusOK, response)
//}
//
//func (w *Wrapper) Json(response *Response) {
//	w.ctx.JSON(http.StatusOK, response)
//}

// Success 成功响应
func (w *Wrapper) Success(data any) {
	response := NewResponse()
	response.Data = data
	w.render(http.StatusOK, response)
}

// DefaultError 默认错误，内置错误信息
func (w *Wrapper) DefaultError(code errcode.Code) {
	response := NewNoDataResponse()
	response.ErrCode = int(code)
	response.ErrMsg = ErrCodeInterface(code).String()
	w.render(http.StatusOK, response)
}

// CustomError 自定义错误信息
func (w *Wrapper) CustomError(code int, message string) {
	response := NewNoDataResponse()
	response.ErrCode = code
	response.ErrMsg = message
	w.render(http.StatusOK, response)
}
