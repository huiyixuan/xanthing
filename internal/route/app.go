package route

import (
	"xanthing/internal/utils"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.Use(GinRecovery())
	r.GET("/", func(c *gin.Context) {
		utils.Success(c, gin.H{
			"message": "Hello, welcome to x anything",
		})
	})

	// 设置微信相关路由
	SetWeixinRoutes(r)

	// 设置用户相关路由
	SetUserRoutes(r)
}
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.InternalError(c, "internal error: "+err.(string))
				c.Abort()
			}
		}()
		c.Next()
	}
}
