package model

type WeixinUser struct {
	ID            int64  `json:"id"`
	Openid        string `json:"openid"`
	WeixinAccount string `json:"weixin_account"`
	AddTime       int64  `json:"add_time"`
	UpdateTime    int64  `json:"update_time"`
	IsSubscribe   int    `json:"is_subscribe"`
}
