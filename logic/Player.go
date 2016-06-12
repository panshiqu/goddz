package logic

import (
	"log"
)

// Player 玩家
type Player struct {
	progress int
}

// GetProgress 获取进度
func (p *Player) GetProgress() int {
	return p.progress
}

// OnEvent 事件到来
func (p *Player) OnEvent(message string) {
	log.Println(message)
}
