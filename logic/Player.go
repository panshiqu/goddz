package logic

import "github.com/panshiqu/goddz/wechat"

// Player 玩家
type Player struct {
	openid string // 编号
	game   Game   // 游戏
	cnt    int    // 计数
}

// GetOpenID 获取微信编号
func (p *Player) GetOpenID() string {
	return p.openid
}

// SetOpenID 设置微信编号
func (p *Player) SetOpenID(v string) {
	p.openid = v
}

// GetCnt 获取计数
func (p *Player) GetCnt() int {
	return p.cnt
}

// SetCnt 设置计数
func (p *Player) SetCnt(v int) {
	p.cnt = v
}

// OnEvent 事件到来
func (p *Player) OnEvent(message string) {
	// 计数
	p.cnt++

	switch message {
	// 过河1
	case "1001":
		p.game = new(Game1001)

	// 过河2
	case "1002":
		p.game = new(Game1002)

	// 过河3
	case "1003":
		p.game = new(Game1003)

	// 过桥
	case "1004":
		p.game = new(Game1004)

	default:
		// 校验
		if p.game == nil {
			wechat.PushTextMessage(p.openid, "请先选择游戏")
			return
		}

		// 游戏事件
		wechat.PushTextMessage(p.openid, p.game.OnGameEvent(message))
		return
	}

	// 描述
	wechat.PushTextMessage(p.openid, p.game.Description())

	// 游戏开始
	wechat.PushTextMessage(p.openid, p.game.OnGameStart())
}
