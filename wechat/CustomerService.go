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
	customerServicePostURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

// CustomerServiceText 客服消息
type CustomerServiceText struct {
	ToUser  string             `json:"touser"`
	MsgType string             `json:"msgtype"`
	Text    TextMessageContent `json:"text"`
}

// CustomerServiceImage 客服消息
type CustomerServiceImage struct {
	ToUser  string              `json:"touser"`
	MsgType string              `json:"msgtype"`
	Image   ImageMessageContent `json:"image"`
}

// TextMessageContent 文本消息内容
type TextMessageContent struct {
	Content string `json:"content"`
}

// ImageMessageContent 图片消息内容
type ImageMessageContent struct {
	MediaID string `json:"media_id"`
}

// PushTextMessage 推送文本消息
func PushTextMessage(user string, message string) {
	csm := &CustomerServiceText{
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

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll failed ", err)
	}

	log.Println("#Send:", csm, string(body))
}

// PushImageMessage 推送图片消息
func PushImageMessage(user string, message string) {
	csm := &CustomerServiceImage{
		ToUser:  user,
		MsgType: "image",
		Image:   ImageMessageContent{MediaID: message},
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

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll failed ", err)
	}

	log.Println("#Send:", csm, string(body))
}
