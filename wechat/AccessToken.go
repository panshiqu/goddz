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
	appID               = "wx45361df884eeaf8d"
	appSecret           = "07ac79d826d61af202659ffc2f726aa1"
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
	if err != nil {
		log.Println("http.Get failed ", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("response.StatusCode error")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ioutil.ReadAll failed ", err)
		return
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := &AccessTokenResponse{}
		if err := json.Unmarshal(body, atr); err != nil {
			log.Println("json.Unmarshal failed ", err)
			return
		}

		a.at = atr.AccessToken
		log.Println("AccessToken Refresh ok")
	} else {
		ater := &AccessTokenErrorResponse{}
		if err := json.Unmarshal(body, ater); err != nil {
			log.Println("json.Unmarshal failed ", err)
			return
		}

		log.Println("AccessToken Refresh failed ", ater.Errcode, ater.Errmsg)
	}
}

// ATIns 单例模式
func ATIns() *AccessToken {
	if ins == nil {
		ins = new(AccessToken)
		ins.Refresh()
	}

	return ins
}
