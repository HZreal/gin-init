package response

/**
 * @Author nico
 * @Date 2024-11-17
 * @File: errCode.go
 * @Description:
 */

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ParamsError          = ErrorCode{904030, "请求参数错误"}
	TooManyRequestsError = ErrorCode{904290, "请求频率过多"}
	NoAuthError          = ErrorCode{904010, "未认证!"}
	OperationError       = ErrorCode{905011, "操作错误"}
	UnKnownError         = ErrorCode{999999, "未知错误!"}
)
