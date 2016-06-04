package wechat

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	customerServicePostURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

// CustomerServiceMessage 客服消息
type CustomerServiceMessage struct {
	ToUser  string             `json:"touser"`
	MsgType string             `json:"msgtype"`
	Text    TextMessageContent `json:"text"`
}

// TextMessageContent 文本消息内容
type TextMessageContent struct {
	Content string `json:"content"`
}

// PushTextMessage 推送文本消息
func PushTextMessage(user string, message string) {
	csm := &CustomerServiceMessage{
		ToUser:  user,
		MsgType: "text",
		Text:    TextMessageContent{Content: message},
	}

	body, err := json.MarshalIndent(csm, " ", "  ")
	if err != nil {
		log.Fatal("json.MarshalIndent failed ", err)
	}

	req, err := http.NewRequest("POST", strings.Join([]string{customerServicePostURL, "?access_token=", ATIns().GetAT()}, ""), bytes.NewReader(body))
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
