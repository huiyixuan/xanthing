package model

type WeixinAccount struct {
	ID          int64  `json:"id"`
	AppID       string `json:"app_id"`
	Account     string `json:"account"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
	AddTime     int64  `json:"add_time"`
	UpdateTime  int64  `json:"update_time"`
}
