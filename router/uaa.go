package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaohaihan/ginWeb/handler"
	"net/http"
)

type uaa struct {
	handler.UaaHandler // 嵌入了 UaaHandler 接口
}

func (u *uaa) addRouter(v1 *gin.RouterGroup) {
	uaa := v1.Group("/uaa")     // 定义 "/gin-web/v1/uaa" 路径分组
	uaa.POST("/login", u.login) // 注册登录接口 POST /gin-web/v1/uaa/login
}

func (u *uaa) login(c *gin.Context) {
	req := &handler.LoginReq{} // 创建一个空的请求结构体
	rsp := &handler.LoginRsp{} // 创建一个空的响应结构体

	err := u.UaaHandler.Login(c, req, rsp) // 调用业务逻辑处理器的 Login 方法
	if err != nil {
		return // 如果处理失败，直接返回（实际项目中应该返回错误响应）
	}

	c.JSON(http.StatusOK, rsp) // 返回 HTTP 200 状态和响应数据
}
