package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	ErrCode string `json:"err_code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func ThrowError(c *gin.Context, appErr *AppError, message string) {
	if message != "" {
		appErr.Message = message
	}
	c.Error(appErr)
	c.Abort()
}

var (
	ErrSystem             = &AppError{http.StatusInternalServerError, "SYSTEM_ERROR", "系统异常，请稍后重试"}
	ErrUserNotFound       = &AppError{http.StatusNotFound, "USER_NOT_FOUND", "用户不存在"}
	ErrInvalidCredentials = &AppError{http.StatusUnauthorized, "INVALID_CREDENTIALS", "认证失败"}
	ErrUnauthorized       = &AppError{http.StatusUnauthorized, "UNAUTHORIZED", "权限不足"}
	ErrInvalidParams      = &AppError{http.StatusUnauthorized, "INVALID_PARAMS", "请求参数错误"}
)
