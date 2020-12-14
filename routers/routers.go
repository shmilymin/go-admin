package routers

import (
	"go-admin/pkg/mid"
	"go-admin/routers/menu"
	"go-admin/routers/role"
	"go-admin/routers/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	setUpConfig(r)
	setUpRouter(r)
	return r
}

// 初始化应用设置
func setUpConfig(r *gin.Engine) {
	// 跨域配置
	r.Use(cors.Default())

	//登录
	r.POST("/login", mid.Auth.LoginHandler)

	// 刷新token
	r.GET("/refresh_token", mid.Auth.RefreshHandler)

}

// 设置路由
func setUpRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(mid.Auth.MiddlewareFunc())
	api.Use(mid.CasbinMid())
	{
		user.RegisterRouter(api.Group("/user"))
		role.RegisterRouter(api.Group("/role"))
		menu.RegisterRouter(api.Group("/menu"))
	}
}
