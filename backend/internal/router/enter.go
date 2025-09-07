package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"volunteer-team/backend/internal/handler"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/manager"
)

// RunServer 启动服务器 路由层
func RunServer() {
	r, err := listen()
	if err != nil {
		panic(err.Error())
	}
	r.Run(fmt.Sprintf("%s:%d", configs.Conf.App.Host, configs.Conf.App.Port)) // 启动 Gin 服务器
}

// listen 配置 Gin 服务器
func listen() (*gin.Engine, error) {
	r := gin.Default() // 创建默认的 Gin 引擎
	// 注册全局中间件（例如获取 Trace ID）
	manager.RequestGlobalMiddleware(r)
	//配置静态路由，用于访问上传的文件
	r.Static("/uploads", "uploads")
	// 创建 RouteManager 实例
	routeManager := manager.NewRouteManager(r)
	// 注册各业务路由组的具体路由
	registerRoutes(routeManager)
	return r, nil
}

// registerRoutes 注册各业务路由的具体处理函数
func registerRoutes(routeManager *manager.RouteManager) {

	routeManager.RegisterCommonRoutes(func(rg *gin.RouterGroup) {
		rg.GET("/ping", func(c *gin.Context) {
			response.Response(c, "pong", nil)
		})
	})

	routeManager.RegisterUserRoutes(func(rg *gin.RouterGroup) {
		userHandler := handler.NewUserHandler()
		rg.POST("/login", middleware.BindJsonMiddleware[types.LoginReq], userHandler.LoginHandler)
	})
}
