package wechat

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

const (
	createCustomMenuPostURL = "https://api.weixin.qq.com/cgi-bin/menu/create"
)

var menu = []byte(`{
    "button": [
        {
            "name": "过河",
            "sub_button": [
                {
                    "type": "click",
                    "name": "过河1",
                    "key": "1001"
                },
                {
                    "type": "click",
                    "name": "过河2",
                    "key": "1002"
                },
                {
                    "type": "click",
                    "name": "过河3",
                    "key": "1003"
                }
            ]
        },
        {
            "type": "click",
            "name": "过桥",
            "key": "1004"
        }
    ]
}`)

// CreateCustomMenu 创建自定义菜单
func CreateCustomMenu() {
	req, err := http.NewRequest("POST", strings.Join([]string{createCustomMenuPostURL, "?access_token=", ATIns().GetAT()}, ""), bytes.NewReader(menu))
	if err != nil {
		log.Fatal("http.NewRequest failed ", err)
	}

	req.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("http.Client.Do failed ", err)
	}
}
