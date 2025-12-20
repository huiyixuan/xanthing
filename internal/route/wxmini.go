package route

import (
	"xanthing/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetWxMiniRoutes 设置微信小程序相关路由
func SetWxMiniRoutes(r *gin.Engine) {
	// 创建微信小程序控制器实例（需要传入实际的AppID和AppSecret）
	// 这里使用示例值，实际项目中应该从配置文件中读取
	wxMiniCtl := controller.NewWxMiniCtl("your_app_id", "your_app_secret")

	// 微信小程序路由组
	wxMiniGroup := r.Group("/wxmini")
	{
		// 微信登录
		wxMiniGroup.POST("/wxLogin", wxMiniCtl.WxLogin)

		// 获取用户信息
		//wxMiniGroup.GET("/user/info", wxMiniCtl.GetUserInfo)
		//
		//// 获取access_token
		//wxMiniGroup.GET("/access_token", wxMiniCtl.GetAccessToken)
		//
		//// 获取用户手机号
		//wxMiniGroup.POST("/phone", wxMiniCtl.GetPhoneNumber)
		//
		//// 发送订阅消息
		//wxMiniGroup.POST("/message/send", wxMiniCtl.SendMessage)
		//
		//// 检查session有效性
		//wxMiniGroup.GET("/session/check", wxMiniCtl.CheckSession)

		// 健康检查
		wxMiniGroup.GET("/health", wxMiniCtl.HealthCheck)
	}
}
