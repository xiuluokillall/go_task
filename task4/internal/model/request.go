package model

type UserParam struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

type PostParam struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type CommentParam struct {
	Content string `json:"content" form:"content"`
}
