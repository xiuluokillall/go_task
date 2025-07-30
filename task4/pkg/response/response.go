package response

import (
	"github.com/gin-gonic/gin"
	error2 "githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  msg,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func Error(c *gin.Context, appErr *error2.AppError) {
	c.JSON(http.StatusOK, Response{
		Code: appErr.Code,
		Msg:  appErr.Message,
		Data: nil,
	})
}

func FailStop(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
