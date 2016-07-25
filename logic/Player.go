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
	// 狼羊菜过河
	case "1001":
		p.game = new(Game1001)

	// 三人三鬼过河
	case "1002":
		p.game = new(Game1002)

	// 警犯一家人过河
	case "1003":
		p.game = new(Game1003)

	// 一家人过独木桥
	case "1004":
		p.game = new(Game1004)

	// 指挥电梯上下
	case "1005":
		p.game = new(Game1005)

	// 青蛙直线跳棋
	case "1006":
		p.game = new(Game1006)

	// 哥弟俩均分饮料
	case "1007":
		p.game = new(Game1007)

	// 切分金条发工资
	case "1008":
		p.game = new(Game1008)

	// 把灯全打开
	case "1009":
		p.game = new(Game1009)

	// 智力测试图集
	case "1010":
		//p.game = new(Game1010)

	default:
		// 校验
		if p.game == nil {
			wechat.PushTextMessage(p.openid, "请先选择游戏")
			return
		}

		switch message {
		// 问题背景
		case "qb":
			wechat.PushTextMessage(p.openid, p.game.Background())
			return

		// 操作说明
		case "oi":
			wechat.PushTextMessage(p.openid, p.game.Description())
			return

		// 当前场景
		case "cs":
			wechat.PushTextMessage(p.openid, p.game.GameScene())
			return

		// 获取提示
		case "tip":
			wechat.PushTextMessage(p.openid, p.game.GameTips())
			return
		}

		// 游戏事件
		wechat.PushTextMessage(p.openid, p.game.OnGameEvent(message))
		return
	}

	// 问题背景
	wechat.PushTextMessage(p.openid, p.game.Background())

	// 操作说明
	wechat.PushTextMessage(p.openid, p.game.Description())

	// 游戏开始
	wechat.PushTextMessage(p.openid, p.game.OnGameStart())
}
