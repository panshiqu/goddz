package main

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozillazg/go-pinyin"
	"github.com/panshiqu/goddz/base"
	"github.com/panshiqu/goddz/logic"
	"github.com/panshiqu/goddz/wechat"
)

const (
	token = "panshiqu"
)

// RequestBody 请求
type RequestBody struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string
}

// MenuRequestBody 菜单请求
type MenuRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Event        string
	EventKey     string
}

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

// VoiceRequestBody 语音请求
type VoiceRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MediaId      string
	Format       string
	Recognition  string
	MsgID        int
}

// TextResponseBody 文本响应
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

// CDATAText CDATA
type CDATAText struct {
	Text string `xml:",innerxml"`
}

func value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func makeSignature(timestamp string, nonce string) string {
	slice := []string{token, timestamp, nonce}
	sort.Strings(slice)

	sha := sha1.New()
	io.WriteString(sha, strings.Join(slice, ""))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func procRequest(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("http.Request.ParseForm failed ", err)
		return
	}

	signature := strings.Join(r.Form["signature"], "")
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")

	if signature != makeSignature(timestamp, nonce) {
		http.Redirect(w, r, "http://www.iplaygame.com.cn:8080"+r.URL.RequestURI(), http.StatusFound)
		log.Println("Validate failed")
		log.Println(r)
		return
	}

	// 接入验证
	//log.Println("Validate ok")
	//fmt.Fprintf(w, strings.Join(r.Form["echostr"], ""))

	// 解析消息
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ioutil.ReadAll failed ", err)
			return
		}

		requestType := &RequestBody{}
		if err := xml.Unmarshal(body, requestType); err != nil {
			log.Println("xml.Unmarshal failed ", err)
			return
		}

		msgType := strings.ToUpper(requestType.MsgType)
		if msgType == "TEXT" {
			requestBody := &TextRequestBody{}
			if err := xml.Unmarshal(body, requestBody); err != nil {
				log.Println("xml.Unmarshal failed ", err)
				return
			}

			log.Println("#Recv:", requestBody)

			if strings.ToLower(requestBody.Content) == "id" {
				wechat.PushTextMessage(requestBody.FromUserName, requestBody.FromUserName)
				return
			}

			if len(requestBody.Content) > 12 && strings.ToLower(requestBody.Content[:4]) == "bind" {
				ss := strings.Split(requestBody.Content, " ")
				if len(ss) != 3 {
					wechat.PushTextMessage(requestBody.FromUserName, "绑定格式错误")
					return
				}
				resp, err := http.Get(fmt.Sprintf(`http://localhost:8080/bind?id=%s&alias=%s&password=%s&admin=true`, requestBody.FromUserName, ss[1], ss[2]))
				if err != nil {
					wechat.PushTextMessage(requestBody.FromUserName, err.Error())
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					wechat.PushTextMessage(requestBody.FromUserName, err.Error())
					return
				}
				wechat.PushTextMessage(requestBody.FromUserName, string(body))
				return
			}

			go logic.PIns().OnEvent(requestBody.FromUserName, requestBody.Content)

			responseBody := &TextResponseBody{
				ToUserName:   value2CDATA(requestBody.FromUserName),
				FromUserName: value2CDATA(requestBody.ToUserName),
				CreateTime:   time.Duration(time.Now().Unix()),
				MsgType:      value2CDATA("text"),
				Content:      value2CDATA(""),
			}

			//log.Println(responseBody)

			text, err := xml.MarshalIndent(responseBody, " ", "  ")
			if err != nil {
				log.Println("xml.MarshalIndent failed ", err)
				return
			}

			w.Header().Set("Content-Type", "text/xml")
			fmt.Fprintf(w, "success")
			_ = string(text)
		} else if msgType == "EVENT" {
			requestBody := &MenuRequestBody{}
			if err := xml.Unmarshal(body, requestBody); err != nil {
				log.Println("xml.Unmarshal failed ", err)
				return
			}

			log.Println("#Recv:", requestBody)

			event := strings.ToUpper(requestBody.Event)
			if event == "SCAN" {
				go logic.PIns().OnScan(requestBody.FromUserName, requestBody.EventKey)
			} else if event == "CLICK" {
				go logic.PIns().OnEvent(requestBody.FromUserName, requestBody.EventKey)
			} else if event == "SUBSCRIBE" {
				go logic.PIns().OnSubscribe(requestBody.FromUserName, requestBody.EventKey)
			} else if event == "UNSUBSCRIBE" {
				go logic.PIns().OnUnSubscribe(requestBody.FromUserName)
			}

			w.Header().Set("Content-Type", "text/xml")
			fmt.Fprintf(w, "success")
		} else if msgType == "VOICE" {
			requestBody := &VoiceRequestBody{}
			if err := xml.Unmarshal(body, requestBody); err != nil {
				log.Println("xml.Unmarshal failed ", err)
				return
			}

			log.Println("#Recv:", requestBody)

			reco := strings.TrimRightFunc(requestBody.Recognition, func(v rune) bool {
				if v == '！' {
					return true
				}

				return false
			})

			var trans string
			a := pinyin.NewArgs()
			for _, v := range pinyin.Pinyin(reco, a) {
				trans += strings.Join(v, "")
			}

			wechat.PushTextMessage(requestBody.FromUserName, "识别结果："+trans)

			go logic.PIns().OnEvent(requestBody.FromUserName, trans)

			w.Header().Set("Content-Type", "text/xml")
			fmt.Fprintf(w, "success")
		}
	}
}

func main() {
	log.Println("start")

	// DB
	if !logic.PIns().InitDB() {
		log.Fatal("logic.PIns.InitDB failed")
	}

	// SSDB
	// if !logic.PIns().InitSSDB() {
	// 	log.Fatal("logic.PIns.InitSSDB failed")
	// }

	// 定期刷新
	base.TMIns().AddTimer(logic.RefreshTimer, time.Duration(time.Hour), true, nil)

	// 通知管理
	wechat.PushTextMessage(logic.AdminOpenID, "服务器启动成功")

	// 创建菜单
	//wechat.CreateCustomMenu()

	// 开启服务
	http.HandleFunc("/", procRequest)
	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		wechat.PushTextMessage(r.FormValue("id"), r.FormValue("message"))
	})
	http.HandleFunc("/favicon.ico", func(http.ResponseWriter, *http.Request) {})
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("http.ListenAndServe failed ", err)
	}
}
