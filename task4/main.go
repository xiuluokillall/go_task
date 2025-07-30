package main

import (
	"githubgithub.com/xiuluokillall/go_task/task4/internal/config"
)

func main() {
	//db := dao.InitMysqlDb()
	// 自动迁移模型
	//db.AutoMigrate(&dao.User{}, &dao.Post{}, &dao.Comment{})
	config.InitConfig("config/config.yaml")
}
