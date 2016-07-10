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
	game     *Game1001
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
		p.key = "1#"
		p.mine = []string{"1", "2"}
		p.rival = []string{"1", "2"}
	case 2:
		p.key = "2#"
		p.mine = []string{"1", "2", "3"}
		p.rival = []string{"1", "2", "3"}
	case 3:
		p.key = "3#"
		p.mine = []string{"1", "2", "3", "4"}
		p.rival = []string{"1", "2", "3", "4"}
	case 4:
		p.key = "4#"
		p.mine = []string{"1", "2", "3", "4", "5"}
		p.rival = []string{"1", "2", "3", "4", "5"}
	case 5:
		p.key = "5#"
		p.mine = []string{"1", "2", "3", "4", "5", "6"}
		p.rival = []string{"1", "2", "3", "4", "5", "6"}
	case 6:
		p.key = "6#"
		p.mine = []string{"1", "2", "3", "4", "5", "6", "7"}
		p.rival = []string{"1", "2", "3", "4", "5", "6", "7"}
	case 7:
		p.key = "7#"
		p.mine = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
		p.rival = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	case 8:
		p.key = "8#"
		p.mine = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		p.rival = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	}
}

// OnEvent 事件到来
func (p *Player) OnEvent(message string) {
	switch message {
	case "cross river game guide":
		game := new(Game1001)
		wechat.PushTextMessage(p.openid, game.Description())
	case "cross river start game":
		p.game = new(Game1001)
		wechat.PushTextMessage(p.openid, p.game.OnGameStart())
	case "run fast game guide":
		wechat.PushTextMessage(p.openid, "游戏说明：\n你拥有优先出牌的权利\n你与机器人拥有指定手牌\n先出完手牌的玩家赢得本轮胜利")
	case "run fast start game":
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
	case "run fast re start":
		// 初始化
		p.Init()

		// 发送场景
		p.OnEvent("run fast start game")
	default:
		p.game.OnGameEvent(message)
		return

		// 输赢判断
		if len(p.mine) == 1 && p.mine[0] == message {
			wechat.PushTextMessage(p.openid, "恭喜你，你赢了")

			// 递增进度
			p.progress++

			// 初始化
			p.Init()

			// 发送场景
			p.OnEvent("run fast start game")

			return
		}

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
			p.key = p.key[:len(p.key)-1]
			return
		}

		// 获取机器操作
		re, err := c.Qget(p.key, PIns().Random().Int63n(size))
		if err != nil {
			log.Fatal("gossdb.Qget ", err)
		}

		// 关键字
		p.key += re.String()

		// 通知用户机器出牌
		wechat.PushTextMessage(p.openid, re.String())

		// 更新场景
		for k, v := range p.mine {
			if v == message {
				p.mine = append(p.mine[:k], p.mine[k+1:]...)
			}
		}
		for k, v := range p.rival {
			if v == re.String() {
				p.rival = append(p.rival[:k], p.rival[k+1:]...)
			}
		}

		// 输赢判断
		if len(p.rival) == 0 {
			wechat.PushTextMessage(p.openid, "很遗憾，你输了")
			return
		}

		// 通知场景
		p.OnEvent("run fast start game")
	}
}
