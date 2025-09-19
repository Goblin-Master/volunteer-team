package router

import (
	"fmt"
	"volunteer-team/backend/internal/handler"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/manager"

	"github.com/gin-gonic/gin"
)

// RunServer 启动服务器 路由层
func RunServer() {
	r, err := listen()
	if err != nil {
		panic(err.Error())
		return
	}
	err = r.Run(fmt.Sprintf("%s:%d", configs.Conf.App.Host, configs.Conf.App.Port)) // 启动 Gin 服务器
	if err != nil {
		panic(err.Error())
	}
}

// listen 配置 Gin 服务器
func listen() (*gin.Engine, error) {
	r := gin.Default() // 创建默认的 Gin 引擎
	// 注册全局中间件（例如获取 Trace ID）
	manager.RequestGlobalMiddleware(r)
	//配置静态路由，用于访问上传的文件
	r.Static("/uploads", configs.Conf.App.Uploads)
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
		rg.POST("/login", middleware.BindJsonMiddleware[types.LoginReq], userHandler.Login)
		rg.POST("/register", middleware.BindJsonMiddleware[types.RegisterReq], userHandler.Register)
		rg.POST("/code/login", middleware.BindJsonMiddleware[types.GetCodeReq], userHandler.GetLoginCode)
		rg.POST("/code/register", middleware.BindJsonMiddleware[types.GetCodeReq], userHandler.GetRegisterCode)
		rg.POST("/code/reset", middleware.BindJsonMiddleware[types.GetCodeReq], userHandler.GetResetCode)
		rg.POST("/resetPassword", middleware.BindJsonMiddleware[types.ResetPasswordReq], userHandler.ResetPassword)
		rg.POST("/updateAvatar", middleware.Authentication(jwtx.COMMON_USER), userHandler.UpdateAvatar)
	})

	routeManager.RegisterOrderRoutes(func(rg *gin.RouterGroup) {
		orderHandler := handler.NewOrderHandler()
		rg.POST("/create", middleware.Authentication(jwtx.COMMON_USER), middleware.BindJsonMiddleware[types.CreateOrderReq], orderHandler.CreateOrder)
		rg.GET("/list", middleware.Authentication(jwtx.COMMON_USER), orderHandler.GetOrderList)
		rg.GET("/detail", middleware.Authentication(jwtx.COMMON_USER), middleware.BindQueryMiddleware[types.OrderDetailReq], orderHandler.GetOrderDetail)
		rg.PUT("/finish", middleware.Authentication(jwtx.INTERNAL_USER), middleware.BindQueryMiddleware[types.FinishOrderReq], orderHandler.FinishOrder)
	})

	routeManager.RegisterSummaryRoutes(func(rg *gin.RouterGroup) {
		summaryHandler := handler.NewSummaryHandler()
		rg.POST("/create", middleware.Authentication(jwtx.INTERNAL_USER), middleware.BindJsonMiddleware[types.CreateSummaryReq], summaryHandler.CreateSummary)
		rg.GET("/list", middleware.Authentication(jwtx.INTERNAL_USER), summaryHandler.GetSummaryList)
		rg.GET("/detail", middleware.Authentication(jwtx.INTERNAL_USER), middleware.BindQueryMiddleware[types.SummaryDetailReq], summaryHandler.GetSummaryDetail)
		rg.PUT("/update", middleware.Authentication(jwtx.INTERNAL_USER), middleware.BindJsonMiddleware[types.UpdateSummaryReq], summaryHandler.UpdateSummary)
	})
}
