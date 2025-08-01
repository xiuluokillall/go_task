package handler

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/internal/model"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/auth"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/dao"
	error2 "githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/response"
	"githubgithub.com/xiuluokillall/go_task/task4/utils"
)

func Register(c *gin.Context) {
	var userParam model.UserParam
	if err := c.ShouldBindJSON(&userParam); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}

	// 加密密码
	hashedPassword, err := utils.GenerateFromPassword(userParam.Password)
	if err != nil {
		response.Fail(c, -1, "密码加密失败")
		return
	}
	user := &model.User{
		UserName: userParam.UserName,
		Password: hashedPassword,
	}
	if err := dao.DB.Create(&user).Error; err != nil {
		response.Fail(c, -1, "注册用户失败")
		return
	}

	response.Success(c, nil, "success")
}

func Login(c *gin.Context) {
	var userParam model.UserParam
	if err := c.ShouldBindJSON(&userParam); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}

	loginUser := &model.User{}
	result := dao.DB.Where("username = ?", userParam.UserName).First(loginUser)
	if result.Error != nil || result.RowsAffected == 0 {
		response.Fail(c, -1, "登录失败 用户名密码错误")
	}

	err := utils.CompareHashAndPassword(loginUser.Password, userParam.Password)
	if err != nil {
		response.Fail(c, -1, "登录失败 密码错误")
	}

	claims := auth.MyClaims{
		UserID:   loginUser.ID,
		Username: loginUser.UserName,
	}
	token, tokenErr := auth.GenerateToken(claims)
	if tokenErr != nil {
		response.Fail(c, -1, "生成token出错")
	}

	c.Set(USERID, claims.UserID)

	response.Success(c, token, "success")
}
