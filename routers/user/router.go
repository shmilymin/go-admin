package user

import "github.com/gin-gonic/gin"

// 注册路由
func RegisterRouter(r *gin.RouterGroup) {
	// 新增
	r.POST("", Add)
	// 根据id查找
	r.GET(":id", Get)
	// 修改
	r.PUT("", Update)
	// 查询全部
	r.GET("", List)
	// 删除
	r.DELETE(":id", Delete)
}
