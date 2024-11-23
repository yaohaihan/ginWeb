package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaohaihan/ginWeb/handler"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	root := router.Group("/gin-web")
	v1 := root.Group("/v1")
	uaa := &uaa{handler.NewUaaHandler()}
	uaa.addRouter(v1)
	return router
}
