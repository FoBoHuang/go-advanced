package errors

import "fmt"

// ResponseSuccess 定义响应成功
type ResponseSuccess struct {
	Code    int         // 响应码
	Message string      // 消息
	Data    interface{} // 数据
}

// ResponseError 定义响应错误
type ResponseError struct {
	Code       int    // 错误码
	Message    string // 错误消息
	StatusCode int    // 响应状态码
	ERR        error  // 响应错误
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

// UnWrapResponse 解包响应错误
func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

// WrapResponse 包装响应错误
func WrapResponse(err error, code, statusCode int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		ERR:        err,
		StatusCode: statusCode,
	}
	return res
}

// Wrap400Response 包装错误码为400的响应错误
func Wrap400Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 400, 400, msg, args...)
}

// Wrap500Response 包装错误码为500的响应错误
func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 500, 500, msg, args...)
}

// NewResponse 创建响应错误
func NewResponse(code, statusCode int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		StatusCode: statusCode,
	}
	return res
}

// New400Response 创建状态码为400的响应错误
func New400Response(code int, msg string, args ...interface{}) error {
	return NewResponse(code, 400, msg, args...)
}

// New500Response 创建状态码为500的响应错误
func New500Response(code int, msg string, args ...interface{}) error {
	return NewResponse(code, 500, msg, args...)
}

// New403Response 创建状态码为403的响应错误
func New403Response(code int, msg string, args ...interface{}) error {
	return NewResponse(code, 403, msg, args...)
}
