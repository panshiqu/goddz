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
                    "name": "开始游戏",
                    "key": "cross river start game"
                },
                {
                    "type": "click",
                    "name": "游戏说明",
                    "key": "cross river game guide"
                }
            ]
        },
        {
            "name": "跑得快",
            "sub_button": [
                {
                    "type": "click",
                    "name": "开始游戏",
                    "key": "run fast start game"
                },
                {
                    "type": "click",
                    "name": "重新开始",
                    "key": "run fast re start"
                },
                {
                    "type": "click",
                    "name": "游戏说明",
                    "key": "run fast game guide"
                }
            ]
        },
        {
            "type": "click",
            "name": "联系我们",
            "key": "contact us"
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
