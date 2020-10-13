package errcode

import (
	"fmt"
	"net/http"
)

// 错误处理

type Error struct {
	code		int			`json:"code"`
	msg 		string		`json:"msg"`
	details		[]string	`json:"details"`
}
// 已存在的错误码
var codes = map[int]string{}

// 添加新的错误码
func NewError(code int, msg string) *Error {
	// 错误码已经存在 报错
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个",code))
	}
	// 不存在 添加错误码
	codes[code] = msg
	return &Error{code: code,msg: msg}
}

// 报错
func (e *Error) Error() string {
	return fmt.Sprintf("错误码 %d ，错误信息：%s",e.code,e.msg)
}

// 返回错误码
func (e *Error) Code() int {
	return e.code
}

// 返回错误信息
func (e *Error) Msg() string {
	return e.msg
}

// 返回组合错误信息
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg,args...)
}

// 返回错误明细
func (e *Error) Details() []string {
	return e.details
}

// 修改错误明细
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _,d := range details {
		e.details = append(e.details,d)
	}
	return e
}

// 返回对应 Http 状态码
func (e *Error) StatusCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}