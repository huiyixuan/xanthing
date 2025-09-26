package wechat

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
)

type Official struct {
}

type Message struct {
	FromUserName string `xml:"FromUserName"`
	ToUserName   string `xml:"ToUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
	PicUrl       string `xml:"PicUrl"`
}

func (o *Official) GetAccessToken() {

}

// ParseXml 解析xml
func (o *Official) ParseXml(xmlStr string) (Message, error) {
	var msg Message
	err := xml.Unmarshal([]byte(xmlStr), &msg)
	if err != nil {
		fmt.Println("parse error:", err)
		return msg, err
	}

	fmt.Println(msg)
	return msg, nil
}

func (o *Official) CheckSign(data map[string]string) bool {
	params := []string{
		"xanthing", data["timestamp"], data["nonce"],
	}
	sort.Strings(params)
	signStr := strings.Join(params, "")
	hasher := sha1.New()
	hasher.Write([]byte(signStr))
	computedSignature := hex.EncodeToString(hasher.Sum(nil))
	if computedSignature == data["signature"] {
		return true
	}
	return false
}
