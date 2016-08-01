package wechat

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	createQrcodePostURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
)

var qrcode = []byte(`{
    "action_name": "QR_LIMIT_STR_SCENE",
    "action_info": {
        "scene": {
            "scene_str": "SuiXian Park"
        }
    }
}`)

// CreateQrcode 创建
func CreateQrcode() {
	req, err := http.NewRequest("POST", strings.Join([]string{createQrcodePostURL, "?access_token=", ATIns().GetAT()}, ""), bytes.NewReader(qrcode))
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
