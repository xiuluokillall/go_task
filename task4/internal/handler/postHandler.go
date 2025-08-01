package handler

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/internal/model"
	error2 "githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/response"
)

func CreatePost(c *gin.Context) {
	var postParam model.PostParam
	if err := c.ShouldBindJSON(&postParam); err != nil {
		response.Error(c, error2.ErrInvalidParams)
	}

	userId := c.MustGet(USERID).(uint)
	post := model.Post{
		UserID:  userId,
		Title:   postParam.Title,
		Content: postParam.Content,
	}

	result := db.Create(&post)
	if result.Error != nil || result.RowsAffected == 0 {
		response.Fail(c, -1, "创建文章失败")
		return
	}

	response.Success(c, nil, "success")
}
