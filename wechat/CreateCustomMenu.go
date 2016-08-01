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
            "name": "游戏列表",
            "sub_button": [
                {
                    "type": "click",
                    "name": "狼羊菜过河",
                    "key": "1001"
                },
                {
                    "type": "click",
                    "name": "三人三鬼过河",
                    "key": "1002"
                },
                {
                    "type": "click",
                    "name": "警犯一家人过河",
                    "key": "1003"
                },
                {
                    "type": "click",
                    "name": "一家人过独木桥",
                    "key": "1004"
                },
                {
                    "type": "click",
                    "name": "指挥电梯上下",
                    "key": "1005"
                }
            ]
        },
        {
            "name": "游戏列表",
            "sub_button": [
                {
                    "type": "click",
                    "name": "青蛙直线跳棋",
                    "key": "1006"
                },
                {
                    "type": "click",
                    "name": "哥弟俩均分饮料",
                    "key": "1007"
                },
                {
                    "type": "click",
                    "name": "切分金条发工资",
                    "key": "1008"
                },
                {
                    "type": "click",
                    "name": "把灯全打开",
                    "key": "1009"
                },
                {
                    "type": "click",
                    "name": "智力测试图集",
                    "key": "1010"
                }
            ]
        },
        {
            "name": "游戏信息",
            "sub_button": [
                {
                    "type": "click",
                    "name": "问题背景",
                    "key": "qb"
                },
                {
                    "type": "click",
                    "name": "操作说明",
                    "key": "oi"
                },
                {
                    "type": "click",
                    "name": "当前场景",
                    "key": "cs"
                },
                {
                    "type": "click",
                    "name": "获取提示",
                    "key": "tip"
                },
                {
                    "type": "click",
                    "name": "游戏攻略",
                    "key": "strategy"
                }
            ]
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

	if err != nil {
		log.Fatal("http.Client.Do failed ", err)
	}

	resp.Body.Close()
}
