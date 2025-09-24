package route

import (
	"net/http"
	"xanthing/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.Use(GinRecovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, welcome to x anything",
		})
	})
	weixin := r.Group("/weixin/")
	{
		weixinCtl := controller.WeixinCtl{}
		weixin.Any("/publicCallback", weixinCtl.PublicCallback)
	}
}
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "internal error: " + err.(string),
					"data": map[string]any{},
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
