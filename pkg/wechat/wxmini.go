package wechat

import (
	"encoding/json"
	"fmt"
	curl "xanthing/internal/utils"
)

// WxMini 微信小程序结构体
type WxMini struct {
	AppID     string
	AppSecret string
}

// NewWxMini 创建微信小程序实例
func NewWxMini(appID, appSecret string) *WxMini {
	return &WxMini{
		AppID:     appID,
		AppSecret: appSecret,
	}
}

// Code2SessionResponse 微信登录响应结构
type Code2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// AccessTokenResponse 获取access_token响应结构
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// UserInfoResponse 获取用户信息响应结构
type UserInfoResponse struct {
	OpenID    string `json:"openid"`
	Nickname  string `json:"nickname"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId,omitempty"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

func (w *WxMini) WxLogin(code string) (Code2SessionResponse, error) {

	response := Code2SessionResponse{
		OpenID:     "mock_openid",
		SessionKey: "mock_session_key",
	}

	req := curl.NewRequest()
	req.URL = "https://api.weixin.qq.com/sns/jscode2session"
	req.Method = "GET"
	req.URL += fmt.Sprintf("?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", w.AppID, w.AppSecret, code)

	res, err := req.Send()
	if err != nil {
		return Code2SessionResponse{}, err
	}
	if err := json.Unmarshal(res.Body, &response); err != nil {
		return Code2SessionResponse{}, err
	}
	fmt.Println(string(res.Body))
	return response, nil
}
