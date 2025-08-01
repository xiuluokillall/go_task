package main

import (
	"githubgithub.com/xiuluokillall/go_task/task4/internal/config"
	"githubgithub.com/xiuluokillall/go_task/task4/pkg/dao"
	"githubgithub.com/xiuluokillall/go_task/task4/router"
)

func main() {
	//db := dao.InitMysqlDb()
	// 自动迁移模型
	//db.AutoMigrate(&dao.User{}, &dao.Post{}, &dao.Comment{})
	config.InitConfig("internal/config/config.yaml")
	//初始化数据库
	dao.InitMysqlDb()
	//启动服务
	router := router.InitRouter()
	err := router.Run(":" + config.GetConfig().Server.Port)
	if err != nil {
		panic(err)
	}
}
