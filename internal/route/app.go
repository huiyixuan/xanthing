package route

import (
	"xanthing/internal/service"
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
	r.GET("/testRedis", func(c *gin.Context) {
		rdb := service.Rdb
		res := rdb.Ping()
		if res.Err() == nil {
			utils.Success(c, gin.H{
				"message": "redis connected",
			})
		} else {
			utils.InternalError(c, "redis connect failed: "+res.Err().Error())
		}
	})

	r.GET("/testMysql", func(c *gin.Context) {
		db, err := service.GetDb("mysql")
		if err != nil {
			utils.InternalError(c, "mysql connect failed: "+err.Error())
			return
		}
		sqlDB, err := db.DB()
		if err != nil {
			utils.InternalError(c, "mysql connect failed: "+err.Error())
			return
		}
		err = sqlDB.Ping()
		if err != nil {
			utils.InternalError(c, "mysql connect failed: "+err.Error())
			return
		}
		utils.Success(c, gin.H{
			"message": "mysql connected",
		})
	})
	// 设置微信相关路由
	SetWeixinRoutes(r)

	// 设置用户相关路由
	SetUserRoutes(r)

	// 设置微信小程序相关路由
	SetWxMiniRoutes(r)
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
