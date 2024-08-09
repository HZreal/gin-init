package common

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessWithData(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func SuccessWithoutData() *Response {
	return &Response{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
}

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var ParamsError = ErrorCode{904030, "请求参数错误"}
var NoAuth = ErrorCode{904010, "未认证!"}

var OperationError = ErrorCode{905011, "操作错误"}
var UKnownError = ErrorCode{999999, "未知错误!"}

func Failed(err ErrorCode) *Response {
	return &Response{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
