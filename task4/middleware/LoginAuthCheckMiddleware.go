package middleware

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/auth"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"strings"
)

// 验证token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			error.ThrowError(c, error.ErrInvalidCredentials, "未提供认证令牌")
			return
		}

		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			error.ThrowError(c, error.ErrInvalidCredentials, "token格式有问题")
		}

		myClaims, err := auth.ParseToken(tokenString)
		if err != nil {
			error.ThrowError(c, error.ErrInvalidCredentials, "无效的令牌："+err.Error())
			return
		}

		c.Set("userID", myClaims.UserID)
		c.Next()
	}
}
