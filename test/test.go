package main

import (
	"bytes"
	"encoding/xml"
	"log"
	"net/http"
	"time"
)

// TextRequestBody 文本请求
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgID        int
}

func main() {
	trb := &TextRequestBody{
		ToUserName:   "touser",
		FromUserName: "oilXLwUJYuhN9-ml2aq2yIMZFByo",
		CreateTime:   1465720802,
		MsgType:      "text",
		Content:      "run fast",
		MsgID:        1,
	}

	body, err := xml.MarshalIndent(trb, " ", "  ")
	if err != nil {
		log.Fatal("xml.MarshalIndent failed ", err)
	}

	url := "http://127.0.0.1/?signature=6a6bc09d12eea078b0d7e28a765b375ffd422f13&timestamp=1465720802&nonce=1130187312&openid=oilXLwUJYuhN9-ml2aq2yIMZFByo"
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Fatal("http.NewRequest failed ", err)
	}

	req.Header.Set("Content-Type", "text/xml; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("http.Client.Do failed ", err)
	}
}
