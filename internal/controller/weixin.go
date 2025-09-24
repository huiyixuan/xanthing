package controller

import (
	"encoding/json"
	"io"
	"xanthing/internal/service"
	"xanthing/pkg/wechat"

	"github.com/gin-gonic/gin"
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
			if err != nil {
				c.String(400, "error:"+err.Error())
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

			c.String(200, "success")

		} else {
			c.String(200, "hello world")
		}
	}
	return
}
