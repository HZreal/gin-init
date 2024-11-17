package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SuccessWithData 成功且返回数据
func SuccessWithData(c *gin.Context, data interface{}) {
	response := &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}

// SuccessWithoutData 成功但不返回数据
func SuccessWithoutData(c *gin.Context) {
	SuccessWithData(c, nil)
}

// Failed 失败返回错误对象
func Failed(c *gin.Context, err ErrorCode) {
	response := &Response{
		Code: err.Code,
		Msg:  err.Msg,
		Data: nil,
	}

	c.JSON(http.StatusOK, response)

}

// FailedWithMsg 失败返回错误字符串
func FailedWithMsg(c *gin.Context, msg string) {
	response := &Response{
		Code: 999999,
		Msg:  msg,
		Data: nil,
	}

	c.JSON(http.StatusOK, response)
}
