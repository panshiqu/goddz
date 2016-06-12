package logic

import (
	"log"
	"strings"

	"github.com/panshiqu/goddz/wechat"
)

// Player 玩家
type Player struct {
	progress int
	openid   string
	key      string
	mine     []string
	rival    []string
}

// GetProgress 获取进度
func (p *Player) GetProgress() int {
	return p.progress
}

// SetProgress 设置进度
func (p *Player) SetProgress(v int) {
	p.progress = v
}

// GetOpenID 获取微信编号
func (p *Player) GetOpenID() string {
	return p.openid
}

// SetOpenID 设置微信编号
func (p *Player) SetOpenID(v string) {
	p.openid = v
}

// Init 初始化
func (p *Player) Init() {
	switch p.progress {
	case 1:
		p.mine = []string{"1", "2"}
		p.rival = []string{"1", "2"}
	case 2:
		p.mine = []string{"1", "2", "3"}
		p.rival = []string{"1", "2", "3"}
	}
}

// OnEvent 事件到来
func (p *Player) OnEvent(message string) {
	switch message {
	case "run fast":
		// 初始化游戏
		if p.progress == 0 {
			p.progress = 1
			p.Init()
		}

		// 构造场景
		mine := "我的：" + strings.Join(p.mine, ", ")
		rival := "机器：" + strings.Join(p.rival, ", ")

		// 发送场景
		wechat.PushTextMessage(p.openid, strings.Join([]string{mine, rival, "请出牌..."}, "\n"))
	default:
		// 关键字
		p.key += message

		// 获取连接
		c, err := PIns().SsdbPool().NewClient()
		if err != nil {
			log.Fatal("gossdb.NewClient ", err)
		}

		// 释放资源
		defer c.Close()

		// 响应数量
		size, err := c.Qsize(p.key)
		if err != nil {
			log.Fatal("gossdb.Qsize ", err)
		}

		// 非法输入
		if size == 0 {
			wechat.PushTextMessage(p.openid, "非法输入，请出牌...")
			return
		}

		// 获取机器操作
		re, err := c.Qget(p.key, PIns().Random().Int63n(size))
		if err != nil {
			log.Fatal("gossdb.Qget ", err)
		}

		// 通知用户机器出牌
		wechat.PushTextMessage(p.openid, re.String())

		// 通知场景
		p.OnEvent("run fast")
	}
}
