package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}
