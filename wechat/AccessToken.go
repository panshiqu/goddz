package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	appID               = "wx8c190a9ee7a26787"
	appSecret           = "a9caec35a7a53b653d7994ade9a70292"
	accessTokenFetchURL = "https://api.weixin.qq.com/cgi-bin/token"
)

// AccessTokenResponse 成功
type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

// AccessTokenErrorResponse 失败
type AccessTokenErrorResponse struct {
	Errcode float64 `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

// AccessToken ...
type AccessToken struct {
	at string
}

// 实例
var ins *AccessToken

// GetAT 获取
func (a *AccessToken) GetAT() string {
	return a.at
}

// Refresh 刷新
func (a *AccessToken) Refresh() {
	request := strings.Join([]string{accessTokenFetchURL,
		"?grant_type=client_credential&appid=",
		appID,
		"&secret=",
		appSecret}, "")

	response, err := http.Get(request)
	defer response.Body.Close()

	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatal("http.Get failed ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll failed ", err)
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := &AccessTokenResponse{}
		if err := json.Unmarshal(body, atr); err != nil {
			log.Fatal("json.Unmarshal failed ", err)
		}

		a.at = atr.AccessToken
		log.Println("AccessToken Refresh ok")
	} else {
		ater := &AccessTokenErrorResponse{}
		if err := json.Unmarshal(body, ater); err != nil {
			log.Fatal("json.Unmarshal failed ", err)
		}

		log.Println("AccessToken Refresh failed ", ater.Errcode, ater.Errmsg)
	}
}

// ATIns 单例模式
func ATIns() *AccessToken {
	if ins == nil {
		ins = new(AccessToken)
	}

	return ins
}
