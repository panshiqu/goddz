package logic

import "strings"

const (
	// Cop 警察
	Cop int = 1

	// Pri 犯人
	Pri int = 2

	// Dad 爸爸
	Dad int = 3

	// Mom 妈妈
	Mom int = 4

	// Son 儿子
	Son int = 5

	// Dau 女儿
	Dau int = 6
)

// Game1003 游戏
type Game1003 struct {
	left    []int          // 左岸
	right   []int          // 右岸
	carry   []int          // 携带
	image   string         // 图片
	voice   map[string]KV  // 语音
	mapping map[int]string // 映射
	name    map[string]int // 名称
	side    bool           // 位置
	cap     int            // 容量
}

// GetID 获取编号
func (g *Game1003) GetID() int {
	return 1003
}

// IsSucceed 是否成功
func (g *Game1003) IsSucceed() bool {
	if len(g.left) == 0 && len(g.carry) == 0 {
		return true
	}

	return false
}

// Background 背景
func (g *Game1003) Background() string {
	return `一家六口人，爸爸、妈妈、两个儿子、两个女儿在旅行途中迷路，幸好遇见一名警察正在押解一名罪犯，无奈只能选择与警察同行寻找回家的路。现在他们需要通过一条河流，你能帮帮他们吗？
注意：船只能同时承载两个货物
注意：只有爸爸、妈妈、警察可以开船
注意：爸爸不在的时候，妈妈便会教训儿子
注意：妈妈不在的时候，爸爸便会教训女儿
注意：警察不在的时候，罪犯会伤害一家六口`
}

// Description 描述
func (g *Game1003) Description() string {
	return `操作：装(1)、卸(2)、过河(3)
货物：警察(1)、罪犯(2)、爸爸(3)、妈妈(4)、儿子(5)、女儿(6)
货物支持简写：警(1)、犯(2)、爸(3)、妈(4)、儿(5)、女(6)

操作和货物之间用点号或空格分隔
装警察上船请输入：装。警察 或 1.1
卸警察下船请输入：卸。警察 或 2.1
过河请输入：过河 或 3

支持语音识别，请说普通话
装警察上船请发送语音：装警察 或 装警
卸警察下船请发送语音：卸警察 或 卸警
过河请发送语音：过河`
}

// OnGameEvent 游戏事件
func (g *Game1003) OnGameEvent(event string) string {
	var cmd int
	var which int
	if kv, ok := g.voice[event]; ok {
		cmd = kv.K
		which = kv.V
	} else {
		var events []string
		if strings.Contains(event, ".") {
			events = strings.Split(event, ".")
		} else if strings.Contains(event, "。") {
			events = strings.Split(event, "。")
		} else if strings.Contains(event, " ") {
			events = strings.Split(event, " ")
		} else {
			events = []string{event}
		}

		cmd, ok = g.name[events[0]]
		if !ok || cmd < Put || cmd > Go {
			return "非法操作"
		}

		if len(events) > 1 && cmd != Go {
			if which, ok = g.name[events[1]]; !ok {
				return "非法货物"
			}
		}
	}

	switch cmd {
	case Put:
		if len(g.carry) >= g.cap {
			return "船已超载"
		}

		if g.side {
			pos := Index(g.left, which)
			if pos == -1 {
				return "请装左岸货物"
			}

			g.left = append(g.left[0:pos], g.left[pos+1:]...)
		} else {
			pos := Index(g.right, which)
			if pos == -1 {
				return "请装右岸货物"
			}

			g.right = append(g.right[0:pos], g.right[pos+1:]...)
		}

		g.carry = append(g.carry, which)
	case Get:
		if len(g.carry) <= 0 {
			return "船已为空"
		}

		pos := Index(g.carry, which)
		if pos == -1 {
			return "请卸船上货物"
		}

		if g.side {
			g.left = append(g.left, which)
		} else {
			g.right = append(g.right, which)
		}

		g.carry = append(g.carry[0:pos], g.carry[pos+1:]...)
	case Go:
		if len(g.carry) <= 0 {
			return "无人驾驶"
		}

		if !Contain(g.carry, Cop) && !Contain(g.carry, Dad) && !Contain(g.carry, Mom) {
			return "无人会驾驶船"
		}

		left := g.left
		right := g.right
		side := !g.side

		for _, v := range g.carry {
			if side {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}

		if !Contain(g.left, Cop) && Contain(g.left, Pri) && len(g.left) > 1 {
			return "左岸罪犯伤害家人"
		}

		if !Contain(g.left, Mom) && Contain(g.left, Dad) && Contain(g.left, Dau) {
			return "左岸爸爸教训女儿"
		}

		if !Contain(g.left, Dad) && Contain(g.left, Mom) && Contain(g.left, Son) {
			return "左岸妈妈教训儿子"
		}

		if !Contain(g.right, Cop) && Contain(g.right, Pri) && len(g.right) > 1 {
			return "右岸罪犯伤害家人"
		}

		if !Contain(g.right, Mom) && Contain(g.right, Dad) && Contain(g.right, Dau) {
			return "右岸爸爸教训女儿"
		}

		if !Contain(g.right, Dad) && Contain(g.right, Mom) && Contain(g.right, Son) {
			return "右岸妈妈教训儿子"
		}

		g.left = left
		g.right = right
		g.side = side
		g.carry = []int{}

		if len(g.left) == 0 {
			return g.GameScene() + "\n\n恭喜过关"
		}
	}

	return g.GameScene()
}

// OnGameStart 游戏开始
func (g *Game1003) OnGameStart() string {
	g.left = []int{Cop, Pri, Dad, Mom, Son, Son, Dau, Dau}
	g.right = []int{}
	g.carry = []int{}
	g.image = "pEnTAPWdIFaIB0fVJT1nv8dBaV-HgMvtc_8ympQi-ec"
	g.voice = map[string]KV{"guohe": {3, 0}, "zhuangjingcha": {1, 1}, "zhuangzuifan": {1, 2}, "zhuangbaba": {1, 3}, "zhuangmama": {1, 4}, "zhuangerzi": {1, 5}, "zhuangnver": {1, 6},
		"zhuangjing": {1, 1}, "zhuangfan": {1, 2}, "zhuangba": {1, 3}, "zhuangma": {1, 4}, "zhuanger": {1, 5}, "zhuangnv": {1, 6},
		"xiejingcha": {2, 1}, "xiezuifan": {2, 2}, "xiebaba": {2, 3}, "xiemama": {2, 4}, "xieerzi": {2, 5}, "xienver": {2, 6},
		"xiejing": {2, 1}, "xiefan": {2, 2}, "xieba": {2, 3}, "xiema": {2, 4}, "xieer": {2, 5}, "xienv": {2, 6}}
	g.mapping = map[int]string{Cop: "警察", Pri: "罪犯", Dad: "爸爸", Mom: "妈妈", Son: "儿子", Dau: "女儿"}
	g.name = map[string]int{"装": Put, "卸": Get, "过河": Go, "警察": Cop, "罪犯": Pri, "爸爸": Dad,
		"妈妈": Mom, "儿子": Son, "女儿": Dau, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6,
		"警": Cop, "犯": Pri, "爸": Dad, "妈": Mom, "儿": Son, "女": Dau}
	g.side = true
	g.cap = 2

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1003) GameImage() string {
	var image string
	if g.image != "" {
		image = g.image
		g.image = ""
	}

	return image
}

// GameScene 游戏场景
func (g *Game1003) GameScene() string {
	scene := "左岸："
	for k, v := range g.left {
		scene += g.mapping[v]
		if k+1 != len(g.left) {
			scene += "、"
		}
	}

	scene += "\n右岸："
	for k, v := range g.right {
		scene += g.mapping[v]
		if k+1 != len(g.right) {
			scene += "、"
		}
	}

	scene += "\n船上："
	for k, v := range g.carry {
		scene += g.mapping[v]
		if k+1 != len(g.carry) {
			scene += "、"
		}
	}

	scene += "\n船在"
	if g.side {
		scene += "左岸"
	} else {
		scene += "右岸"
	}

	return scene
}

// GameTips 提示
func (g *Game1003) GameTips() string {
	return "警犯可是成功的关键，多活动活动他们"
}

// Strategy 攻略
func (g *Game1003) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nv47tmSKO3M7nMI4Fk10b2Vw"
}

// Remind 提醒
func (g *Game1003) Remind() string {
	if len(g.left) == 0 {
		return "您已通关警犯一家人过河，请通过点击菜单或发送指令选择其它游戏继续挑战"
	}

	return "还未通关警犯一家人过河，开动脑筋继续挑战吧，当然您也可以通过点击菜单或发送指令获取提示和攻略\n\n" + g.GameScene()
}
