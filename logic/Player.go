package logic

import (
	"log"
	"wechat"
)

// Player 玩家
type Player struct {
	openid   string
	progress int
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

// OnEvent 事件到来
func (p *Player) OnEvent(message string) {
	switch message {
	case "run fast":
		log.Println("run fast")
		wechat.PushTextMessage(p.openid, "run fast")
	}
}
