package wechat

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	getMediaIDPostURL = "https://api.weixin.qq.com/cgi-bin/material/batchget_material"
)

// 图片（image）、视频（video）、语音（voice）、图文（news）
// 从全部素材的该偏移位置开始返回，0表示从第一个素材返回
var mediaid = []byte(`{
    "type": "news",
    "offset": 0,
    "count": 20
}`)

// GetMediaID 获取
func GetMediaID() {
	req, err := http.NewRequest("POST", strings.Join([]string{getMediaIDPostURL, "?access_token=", ATIns().GetAT()}, ""), bytes.NewReader(mediaid))
	if err != nil {
		log.Fatal("http.NewRequest failed ", err)
	}

	req.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("http.Client.Do failed ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll failed ", err)
	}

	log.Println(string(body))
}
