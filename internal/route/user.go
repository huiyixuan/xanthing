package route

import (
	"xanthing/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetUserRoutes 设置用户相关路由
func SetUserRoutes(r *gin.Engine) {
	user := r.Group("/user/")
	{
		userCtl := controller.UserCtl{}

		// GET请求示例：获取用户信息
		user.GET("info", userCtl.GetUserInfo)

		// POST请求示例：创建用户（支持多种POST格式）
		user.POST("create", userCtl.CreateUser)
	}
}
