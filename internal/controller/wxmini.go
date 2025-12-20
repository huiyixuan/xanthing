package controller

import (
	"fmt"
	"xanthing/internal/utils"
	"xanthing/pkg/wechat"

	"github.com/gin-gonic/gin"
)

// WxMiniCtl 微信小程序控制器
type WxMiniCtl struct {
	wxMini *wechat.WxMini
}

// NewWxMiniCtl 创建微信小程序控制器
func NewWxMiniCtl(appID, appSecret string) *WxMiniCtl {
	return &WxMiniCtl{
		wxMini: wechat.NewWxMini(appID, appSecret),
	}
}

// WxLogin 微信小程序登录
func (w *WxMiniCtl) WxLogin(c *gin.Context) {
	// 获取请求参数
	params, _ := utils.NewRequestParams(c)

	// 获取code参数
	code := params.GetString("code")
	if code == "" {
		utils.BadRequest(c, "缺少code参数")
		return
	}

	// 调用微信登录
	w.wxMini.AppID = "wx38d14e54658adabb"
	w.wxMini.AppSecret = "c24269307582d97a6168f80d4157ad07"
	result, err := w.wxMini.WxLogin(code)
	if err != nil {
		utils.InternalError(c, fmt.Sprintf("微信登录失败: %v", err))
		return
	}

	// 返回登录结果
	utils.Success(c, gin.H{
		"openid":      result.OpenID,
		"session_key": result.SessionKey,
		"unionid":     result.UnionID,
	})
}

// GetUserInfo 获取用户信息
//func (w *WxMiniCtl) GetUserInfo(c *gin.Context) {
//	// 获取请求参数
//	params,_ := utils.NewRequestParams(c)
//
//	// 获取access_token和openid参数
//	accessToken, err := params.GetString("access_token")
//	if err != nil || accessToken == "" {
//		utils.BadRequest(c, "缺少access_token参数")
//		return
//	}
//
//	openID, err := params.GetString("openid")
//	if err != nil || openID == "" {
//		utils.BadRequest(c, "缺少openid参数")
//		return
//	}
//
//	// 获取用户信息
//	userInfo, err := w.wxMini.GetUserInfo(accessToken, openID)
//	if err != nil {
//		utils.InternalError(c, fmt.Sprintf("获取用户信息失败: %v", err))
//		return
//	}
//
//	// 返回用户信息
//	utils.Success(c, gin.H{
//		"openid":    userInfo.OpenID,
//		"nickname":  userInfo.Nickname,
//		"gender":    userInfo.Gender,
//		"city":      userInfo.City,
//		"province":  userInfo.Province,
//		"country":   userInfo.Country,
//		"avatarUrl": userInfo.AvatarURL,
//		"unionId":   userInfo.UnionID,
//	})
//}

//// GetAccessToken 获取access_token
//func (w *WxMiniCtl) GetAccessToken(c *gin.Context) {
//	// 获取access_token
//	result, err := w.wxMini.GetAccessToken()
//	if err != nil {
//		utils.InternalError(c, fmt.Sprintf("获取access_token失败: %v", err))
//		return
//	}
//
//	// 返回access_token
//	utils.Success(c, gin.H{
//		"access_token": result.AccessToken,
//		"expires_in":   result.ExpiresIn,
//	})
//}

// GetPhoneNumber 获取用户手机号
//func (w *WxMiniCtl) GetPhoneNumber(c *gin.Context) {
//	// 获取请求参数
//	params := utils.NewRequestParams(c)
//
//	// 获取access_token和code参数
//	accessToken, err := params.GetString("access_token")
//	if err != nil || accessToken == "" {
//		utils.BadRequest(c, "缺少access_token参数")
//		return
//	}
//
//	code, err := params.GetString("code")
//	if err != nil || code == "" {
//		utils.BadRequest(c, "缺少code参数")
//		return
//	}
//
//	// 获取手机号信息
//	phoneInfo, err := w.wxMini.GetPhoneNumber(accessToken, code)
//	if err != nil {
//		utils.InternalError(c, fmt.Sprintf("获取手机号失败: %v", err))
//		return
//	}
//
//	// 返回手机号信息
//	utils.Success(c, gin.H{
//		"phoneNumber":     phoneInfo.PhoneNumber,
//		"purePhoneNumber": phoneInfo.PurePhoneNumber,
//		"countryCode":     phoneInfo.CountryCode,
//	})
//}
//
//// SendMessage 发送订阅消息
//func (w *WxMiniCtl) SendMessage(c *gin.Context) {
//	// 获取请求参数
//	params := utils.NewRequestParams(c)
//
//	// 获取必需参数
//	accessToken, err := params.GetString("access_token")
//	if err != nil || accessToken == "" {
//		utils.BadRequest(c, "缺少access_token参数")
//		return
//	}
//
//	toUser, err := params.GetString("touser")
//	if err != nil || toUser == "" {
//		utils.BadRequest(c, "缺少touser参数")
//		return
//	}
//
//	templateID, err := params.GetString("template_id")
//	if err != nil || templateID == "" {
//		utils.BadRequest(c, "缺少template_id参数")
//		return
//	}
//
//	// 获取可选参数
//	page, _ := params.GetString("page")
//
//	// 构建消息数据
//	data := make(map[string]map[string]string)
//
//	// 这里可以根据实际需求构建具体的消息数据
//	// 示例：构建一个简单的消息数据
//	data["thing1"] = map[string]string{"value": "测试消息"}
//	data["thing2"] = map[string]string{"value": "这是一条测试消息"}
//
//	// 发送订阅消息
//	err = w.wxMini.SendSubscribeMessage(accessToken, toUser, templateID, page, data)
//	if err != nil {
//		utils.InternalError(c, fmt.Sprintf("发送订阅消息失败: %v", err))
//		return
//	}
//
//	utils.Success(c, gin.H{
//		"message": "订阅消息发送成功",
//	})
//}
//
//// CheckSession 检查session有效性
//func (w *WxMiniCtl) CheckSession(c *gin.Context) {
//	// 获取请求参数
//	params := utils.NewRequestParams(c)
//
//	// 获取必需参数
//	accessToken, err := params.GetString("access_token")
//	if err != nil || accessToken == "" {
//		utils.BadRequest(c, "缺少access_token参数")
//		return
//	}
//
//	openID, err := params.GetString("openid")
//	if err != nil || openID == "" {
//		utils.BadRequest(c, "缺少openid参数")
//		return
//	}
//
//	sessionKey, err := params.GetString("session_key")
//	if err != nil || sessionKey == "" {
//		utils.BadRequest(c, "缺少session_key参数")
//		return
//	}
//
//	// 检查session有效性
//	isValid, err := w.wxMini.CheckSession(accessToken, openID, sessionKey)
//	if err != nil {
//		utils.InternalError(c, fmt.Sprintf("检查session失败: %v", err))
//		return
//	}
//
//	utils.Success(c, gin.H{
//		"is_valid": isValid,
//	})
//}

// HealthCheck 健康检查
func (w *WxMiniCtl) HealthCheck(c *gin.Context) {
	utils.Success(c, gin.H{
		"status":  "ok",
		"service": "微信小程序服务",
	})
}
