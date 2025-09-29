package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
	"xanthing/internal/model"
	"xanthing/internal/service"
	"xanthing/pkg/wechat"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type WeixinCtl struct{}

func (ctl *WeixinCtl) PublicCallback(c *gin.Context) {
	method := c.Request.Method

	params := make(map[string]string)
	_ = c.BindQuery(params)
	official := wechat.Official{}
	if !official.CheckSign(params) {
		c.String(400, "fail")
		return
	}

	if method == "GET" {
		c.String(200, params["echostr"])
	} else {
		contentType := c.GetHeader("Content-Type")
		if contentType == "text/xml" {
			body, err := io.ReadAll(c.Request.Body)
			fmt.Println(string(body))
			if err != nil {
				c.String(400, "error:"+err.Error())
				return
			}
			message, err := official.ParseXml(string(body))
			if err != nil {
				c.String(500, "数据解析失败")
				return
			}

			rdb := service.Rdb
			queue := "x:weixin_official_message"
			msg, err := json.Marshal(message)
			if err == nil {
				rdb.RPush(queue, msg)
			}

			db, _ := service.GetDb("mysql")

			t := time.Unix(cast.ToInt64(message.CreateTime), 0)
			eventDate := t.Format("20060101")

			now := time.Now()
			weixinOfficialMessage := model.WeixinOfficialMessage{
				EventDate:  eventDate,
				FromUser:   message.FromUserName,
				ToUser:     message.ToUserName,
				Content:    message.Content,
				PicUrl:     message.PicUrl,
				AddTime:    now.Unix(),
				UpdateTime: now.Unix(),
				Body:       string(body),
				MsgType:    message.MsgType,
				MsgId:      message.MsgId,
			}
			db.Table("weixin_official_message").Create(&weixinOfficialMessage)

			c.String(200, "success")

		} else {
			c.String(200, "hello world")
		}
	}
	return
}
