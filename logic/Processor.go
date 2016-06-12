package logic

import (
	"wechat"
)

// Processor 处理器
type Processor struct {
	players map[string]*Player
}

// 实例
var ins *Processor

// OnTimer 定时器到期
func (p *Processor) OnTimer(tid int64, param interface{}) {
	switch tid {
	case 1:
		wechat.ATIns().Refresh()
	case 2:
		wechat.PushTextMessage("oilXLwUJYuhN9-ml2aq2yIMZFByo", "定时测试")
	}
}

// OnEvent 事件到来
func (p *Processor) OnEvent(user string, message string) {
	// 查找用户
	player, ok := p.players[user]
	if !ok {
		// 创建用户
		player = new(Player)
		p.players[user] = player
	}

	// 通知事件
	player.OnEvent(message)
}

// PIns 单例模式
func PIns() *Processor {
	if ins == nil {
		ins = new(Processor)
		ins.players = make(map[string]*Player)
	}

	return ins
}
