package main

import (
	"volunteer-team/backend/internal/cmd/flags"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/router"
)

func main() {
	flags.Parse()        //添加一些命令行指令
	configs.LoadConfig() //加载配置
	global.Init()        //初始化全局中间件
	flags.Run()          //可以执行一些命令行指令
	router.RunServer()   //启动服务端
}
