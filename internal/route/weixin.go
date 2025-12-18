package route

import (
	"xanthing/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetWeixinRoutes 设置微信相关路由
func SetWeixinRoutes(r *gin.Engine) {
	weixin := r.Group("/weixin/")
	{
		weixinCtl := controller.WeixinCtl{}
		weixin.Any("/publicCallback", weixinCtl.PublicCallback)
	}

	wxmini := r.Group("/wxmini/")
	{
		wxminiCtl := controller.WxminiCtl{}
		wxmini.POST("/wxLogin", wxminiCtl.WxLogin)
	}

}
