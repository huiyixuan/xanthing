package model

type WeixinOfficialMessage struct {
	ID         int64  `json:"id"`
	EventDate  string `json:"event_date"`
	FromUser   string `json:"from_user"`
	ToUser     string `json:"to_user"`
	Content    string `json:"content"`
	PicUrl     string `json:"pic_url"`
	AddTime    int64  `json:"add_time"`
	UpdateTime int64  `json:"update_time"`
	Body       string `json:"body"`
	MsgType    string `json:"msg_type"`
	MsgId      string `json:"msg_id"`
}
