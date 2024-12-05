package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaohaihan/ginWeb/handler"
)

func InitRoutes() *gin.Engine {
	router := gin.New()                  // 创建一个新的 Gin 路由引擎
	root := router.Group("/gin-web")     // 定义根路径分组 "/gin-web"
	v1 := root.Group("/v1")              // 定义 API 版本分组 "/gin-web/v1"
	uaa := &uaa{handler.newUaaHandler()} // 创建 `uaa` 路由组，并传入处理器
	uaa.addRouter(v1)                    // 调用 `addRouter` 方法注册具体的路由   说白了就是定义了uaa这个handler里面所有的控制逻辑，路由都是从/gin-web/v1开始的
	return router                        // 返回完整的路由引擎
}

//初始化路由引擎：
//
//使用 gin.New() 创建新的 Gin 路由实例，而不是默认的 gin.Default()（避免自动添加中间件，如日志和恢复功能）。
//分组管理路由：
//
///gin-web 是根路径分组，表示 API 的入口。
///v1 是版本分组，用于支持 API 的版本管理（方便将来扩展到 /v2 等）。
//注册子路由：
//
//使用 uaa.addRouter(v1) 将具体的路由规则注册到 /gin-web/v1 路径下。
//依赖注入：
//
//&uaa{handler.NewUaaHandler()} 创建了一个 uaa 结构体，将 handler.UaaHandler（实际是业务逻辑处理器）作为依赖注入到 uaa 中，方便后续使用。
