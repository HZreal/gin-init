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

var (
	ParamsError    = ErrorCode{904030, "请求参数错误"}
	NoAuthError    = ErrorCode{904010, "未认证!"}
	OperationError = ErrorCode{905011, "操作错误"}
	UnKnownError   = ErrorCode{999999, "未知错误!"}
)

func Failed(err ErrorCode) *Response {
	return &Response{
		Code: err.Code,
		Msg:  err.Msg,
		Data: nil,
	}
}

func FailedWithMsg(msg string) *Response {
	return &Response{
		Code: 999999,
		Msg:  msg,
		Data: nil,
	}
}
