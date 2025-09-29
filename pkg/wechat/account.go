package wechat

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
	"xanthing/internal/service"
	curl "xanthing/internal/utils"
)

type Account struct {
	AppID       string
	Account     string
	AppSecret   string
	AccessToken string
}

type Response struct {
	Count int64               `json:"count"`
	Total int64               `json:"total"`
	Data  map[string][]string `json:"data"`
}

func (account *Account) GetAccessToken() (string, error) {
	cacheKey := "access_token"

	rdb := service.Rdb
	cacheData := rdb.Get(cacheKey)
	if cacheData.Val() != "" {
		account.AccessToken = cacheData.Val()
		return cacheData.String(), nil
	}

	fmt.Println("request access token")
	req := curl.NewRequest()
	req.URL = DOMAIN + "/cgi-bin/stable_token"
	req.Method = "POST"
	req.Headers["Content-Type"] = "application/json"

	requestBody := map[string]string{
		"grant_type": "client_credential",
		"appid":      account.AppID,
		"secret":     account.AppSecret,
	}
	body, _ := json.Marshal(requestBody)
	req.Body = body
	resp, err := req.Send()
	if err != nil {
		// 处理错误
		return "", err
	}
	var result map[string]any
	err = json.Unmarshal(resp.Body, &result)

	rdb.Set(cacheKey, result["access_token"].(string), time.Hour*2)
	account.AccessToken = result["access_token"].(string)
	return account.AccessToken, nil
}

func (account *Account) GetUsers(nextOpenid string) (Response, error) {
	var response Response
	req := curl.NewRequest()
	req.URL = DOMAIN + "/cgi-bin/user/get"
	req.Method = "GET"
	query := map[string]string{
		"access_token": account.AccessToken,
		"next_openid":  nextOpenid,
	}

	// map 转query
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}
	encodedQuery := values.Encode()
	parsedURL, _ := url.Parse(req.URL)
	parsedURL.RawQuery = encodedQuery
	req.URL = parsedURL.String()

	resp, err := req.Send()
	if err != nil {
		// 处理错误
		return response, err
	}
	err = json.Unmarshal(resp.Body, &response)
	return response, nil
}
