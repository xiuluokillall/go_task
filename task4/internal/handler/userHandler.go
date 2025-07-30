package handler

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/internal/model"
	error2 "githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/response"
	"githubgithub.com/xiuluokillall/go_task/task4/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}

	// 加密密码
	hashedPassword, err := utils.GenerateFromPassword(user.Password)
	if err != nil {
		response.Error(c, error2.ErrInvalidParams)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
