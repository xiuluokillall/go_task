package handler

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/internal/model"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/dao"
	error2 "githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/response"
	"githubgithub.com/xiuluokillall/go_task/task4/utils"
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
		response.Fail(c, -1, "密码加密失败")
		return
	}
	user.Password = string(hashedPassword)
	db := dao.GetDb()
	if err := db.Create(&user).Error; err != nil {
		response.Fail(c, -1, "注册用户失败")
		return
	}

	response.Success(c, nil, "success")
}
