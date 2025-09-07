package manager

import (
	"github.com/gin-gonic/gin"
	"volunteer-team/backend/internal/infrastructure/middleware"
)

//主要管理路由组和中间件的注册

// PathHandler 是一个用于注册路由组的函数类型
type PathHandler func(rg *gin.RouterGroup)

// Middleware 是一个用于生成中间件的函数类型
type Middleware func() gin.HandlerFunc

// RouteManager 管理不同的路由组，按业务功能分组
type RouteManager struct {
	UserRoutes   *gin.RouterGroup // 登录相关的路由组
	CommonRoutes *gin.RouterGroup //通用功能相关的路由组
}

// NewRouteManager 创建一个新的 RouteManager 实例，包含各业务功能的路由组
func NewRouteManager(router *gin.Engine) *RouteManager {
	return &RouteManager{
		UserRoutes:   router.Group("/api/user"),   // 初始化登录路由组
		CommonRoutes: router.Group("/api/common"), //通用功能相关的路由组
	}
}

// RegisterUserRoutes 注册登录相关的路由处理函数
func (rm *RouteManager) RegisterUserRoutes(handler PathHandler) {
	handler(rm.UserRoutes)
}

// RegisterCommonRoutes通用功能相关的路由组
func (rm *RouteManager) RegisterCommonRoutes(handler PathHandler) {
	handler(rm.CommonRoutes)
}

// RegisterMiddleware 根据组名为对应的路由组注册中间件
// group 参数为 "login"、"profile"、"team"或"Common"，分别对应不同的路由组
func (rm *RouteManager) RegisterMiddleware(group string, middleware Middleware) {
	switch group {
	case "user":
		rm.UserRoutes.Use(middleware())
	case "common":
		rm.CommonRoutes.Use(middleware())
	}
}

// RequestGlobalMiddleware 注册全局中间件，应用于所有路由
func RequestGlobalMiddleware(r *gin.Engine) {
	r.Use(middleware.Cors())
}
