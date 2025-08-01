package handler

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/internal/model"
	error2 "githubgithub.com/xiuluokillall/go_task/task4/pkg/error"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/response"
)

const USERID = "user_id"

func CreateComment(c *gin.Context) {
	var commentParam model.CommentParam
	if err := c.ShouldBindJSON(&commentParam); err != nil {
		response.Error(c, error2.ErrInvalidParams)
	}

	//userId := c.MustGet(USERID).(uint)
	//comment := model.Comment{
	//	Content: commentParam.Content,
	//	UserID:  userId,
	//}
}
