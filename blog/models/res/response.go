package res

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Succes = 0
	Error  = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// 成功
func Ok(data any, msg string, c *gin.Context) {
	Result(Succes, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Succes, data, "成功", c)
}
func OkWithList(count int64, list any, c *gin.Context) {
	OkWithData(ListResponse[any]{count, list}, c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(Succes, map[string]any{}, msg, c)
}
func OkWith(c *gin.Context) {
	Result(Succes, map[string]any{}, "成功", c)
}

// 失败
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

//func FailWithData(data any, c *gin.Context) {
//	Result(Error, data, "失败", c)
//}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithErr(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMsg(msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)

}
