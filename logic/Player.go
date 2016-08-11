package logic

import (
	"database/sql"
	"log"
	"time"

	"github.com/panshiqu/goddz/wechat"
)

// Player 玩家
type Player struct {
	remind *time.Timer // 提醒
	openid string      // 编号
	game   Game        // 游戏
	cnt    int         // 计数
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

// Init 初始化
func (p *Player) Init() {
	p.remind = time.AfterFunc(24*time.Hour, func() {
		p.OnRemind()
	})
}

// OnEvent 事件到来
func (p *Player) OnEvent(message string) {
	// 计数
	p.cnt++

	// 重新计时
	p.remind.Reset(24 * time.Hour)

	switch message {
	// 狼羊菜过河
	case "1001", "狼羊菜过河", "langyangcaiguohe":
		p.game = new(Game1001)

	// 三人三鬼过河
	case "1002", "三人三鬼过河", "sanrensanguiguohe":
		p.game = new(Game1002)

	// 警犯一家人过河
	case "1003", "警犯一家人过河", "jingfanyijiarenguohe":
		p.game = new(Game1003)

	// 一家人过独木桥
	case "1004", "一家人过独木桥", "yijiarenguodumuqiao":
		p.game = new(Game1004)

	// 指挥电梯上下
	case "1005", "指挥电梯上下", "zhihuidiantishangxia":
		p.game = new(Game1005)

	// 青蛙直线跳棋
	case "1006", "青蛙直线跳棋", "qingwazhixiantiaoqi":
		p.game = new(Game1006)

	// 哥弟俩均分饮料
	case "1007", "哥弟俩均分饮料", "gediliajunfenyinliao":
		p.game = new(Game1007)

	// 切分金条发工资
	case "1008", "切分金条发工资", "qiefenjintiaofagongzi":
		p.game = new(Game1008)

	// 把灯全打开
	case "1009", "把灯全打开", "badengquandakai":
		p.game = new(Game1009)

	// 智力测试图集
	case "1010", "智力测试图集", "zhiliceshituji":
		p.game = new(Game1010)

	default:
		// 校验
		if p.game == nil {
			wechat.PushTextMessage(p.openid, "请先选择游戏")
			return
		}

		switch message {
		// 问题背景
		case "qb", "问题背景", "wentibeijing":
			wechat.PushTextMessage(p.openid, p.game.Background())
			return

		// 操作说明
		case "oi", "操作说明", "caozuoshuoming":
			wechat.PushTextMessage(p.openid, p.game.Description())
			return

		// 当前场景
		case "cs", "当前场景", "dangqianchangjing":
			wechat.PushTextMessage(p.openid, p.game.GameScene())
			return

		// 获取提示
		case "tip", "获取提示", "huoqutishi":
			wechat.PushTextMessage(p.openid, p.game.GameTips())
			return

		// 游戏攻略
		case "strategy", "游戏攻略", "youxigonglve":
			if p.game.Strategy() == "" {
				wechat.PushTextMessage(p.openid, "没有游戏攻略")
			} else {
				wechat.PushMpnewsMessage(p.openid, p.game.Strategy())
			}

			return
		}

		// 游戏事件
		scene := p.game.OnGameEvent(message)

		// 游戏图片
		if image := p.game.GameImage(); image != "" {
			wechat.PushImageMessage(p.openid, image)
		}

		// 游戏场景
		wechat.PushTextMessage(p.openid, scene)

		// 成功通关
		if p.game.IsSucceed() {
			p.OnSuccess()
		}

		return
	}

	// 问题背景
	wechat.PushTextMessage(p.openid, p.game.Background())

	// 操作说明
	wechat.PushTextMessage(p.openid, p.game.Description())

	// 游戏开始
	scene := p.game.OnGameStart()

	// 游戏图片
	if image := p.game.GameImage(); image != "" {
		wechat.PushImageMessage(p.openid, image)
	}

	// 游戏场景
	wechat.PushTextMessage(p.openid, scene)
}

// OnRemind 触发提醒
func (p *Player) OnRemind() {
	if p.game == nil {
		wechat.PushTextMessage(p.openid, "您已24小时未开始游戏，若不知如何操作请仔细阅读新手引导！")
		wechat.PushMpnewsMessage(p.openid, GuideMpnews)
		return
	}

	// 游戏内提醒
	wechat.PushTextMessage(p.openid, p.game.Remind())
}

// OnSuccess 成功通关
func (p *Player) OnSuccess() {
	var count int
	if err := PIns().GetDB().QueryRow("SELECT PassCount FROM PassStat WHERE OpenID = ?", p.openid).Scan(&count); err == nil {
		// 更新数据
		if _, err := PIns().GetDB().Exec("UPDATE PassStat SET PassCount = PassCount + 1 WHERE OpenID = ? AND GameID = ?", p.openid, p.game.GetID()); err != nil {
			log.Println("UPDATE", err)
		}
	} else if err == sql.ErrNoRows {
		// 插入数据
		if _, err := PIns().GetDB().Exec("INSERT INTO PassStat (OpenID, GameID) VALUES (?, ?)", p.openid, p.game.GetID()); err != nil {
			log.Println("INSERT", err)
		}
	} else {
		log.Println("SELECT", err)
	}
}
