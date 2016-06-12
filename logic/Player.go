package logic

import (
	"strings"
	"wechat"
)

// Player 玩家
type Player struct {
	progress int
	openid   string
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
	}
}
