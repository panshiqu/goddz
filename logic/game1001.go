package logic

import "strings"

const (
	// Cabbage 卷心菜
	Cabbage int = 1

	// Sheep 小羊
	Sheep int = 2

	// Wolf 狼
	Wolf int = 3
)

// Game1001 游戏
type Game1001 struct {
	left    []int          // 左岸
	right   []int          // 右岸
	carry   []int          // 携带
	image   string         // 图片
	voice   map[string]KV  // 语音
	mapping map[int]string // 映射
	name    map[string]int // 名称
	race    map[int]int    // 竞争
	side    bool           // 位置
	cap     int            // 容量
}

// Background 背景
func (g *Game1001) Background() string {
	return `农夫准备把卷心菜、小羊、狼运到河对岸的集市售卖，可是他只有一条仅能同时承载一个货物的小船，而且他不在的时候小羊会吃掉卷心菜、狼会吃掉小羊。假如你是他，你能成功将所有货物安全运到对岸吗？`
}

// Description 描述
func (g *Game1001) Description() string {
	return `操作：装(1)、卸(2)、过河(3)
货物：卷心菜(1)、小羊(2)、狼(3)
货物支持简写：菜(1)、羊(2)

操作和货物之间用点号或空格分隔
装狼上船请输入：装。狼 或 1.3
卸狼下船请输入：卸。狼 或 2.3
过河请输入：过河 或 3

支持语音识别，请说普通话
装卷心菜上船请发送语音：装卷心菜 或 装菜
卸卷心菜下船请发送语音：卸卷心菜 或 卸菜
过河请发送语音：过河`
}

// OnGameEvent 游戏事件
func (g *Game1001) OnGameEvent(event string) string {
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

		if side {
			if ok, k, v := g.RaceDetect(right); !ok {
				return "右岸的" + g.mapping[k] + "被" + g.mapping[v] + "吃掉了"
			}
		} else {
			if ok, k, v := g.RaceDetect(left); !ok {
				return "左岸的" + g.mapping[k] + "被" + g.mapping[v] + "吃掉了"
			}
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

// RaceDetect 竞争检测
func (g *Game1001) RaceDetect(s []int) (bool, int, int) {
	for k, v := range g.race {
		if Contain(s, k) && Contain(s, v) {
			return false, k, v
		}
	}

	return true, 0, 0
}

// OnGameStart 游戏开始
func (g *Game1001) OnGameStart() string {
	g.left = []int{Cabbage, Sheep, Wolf}
	g.right = []int{}
	g.carry = []int{}
	g.image = "pEnTAPWdIFaIB0fVJT1nv3xlwsmjRYzoPEa5JVvsXKY"
	g.voice = map[string]KV{"guohe": {3, 0}, "zhuangjuanxincai": {1, 1}, "zhuangxiaoyang": {1, 2}, "zhuanglang": {1, 3}, "zhuangcai": {1, 1},
		"zhuangyang": {1, 2}, "xiejuanxincai": {2, 1}, "xiexiaoyang": {2, 2}, "xielang": {2, 3}, "xiecai": {2, 1}, "xieyang": {2, 2}}
	g.mapping = map[int]string{Cabbage: "卷心菜", Sheep: "小羊", Wolf: "狼"}
	g.name = map[string]int{"装": Put, "卸": Get, "过河": Go, "卷心菜": Cabbage, "小羊": Sheep, "狼": Wolf,
		"1": 1, "2": 2, "3": 3, "菜": Cabbage, "羊": Sheep}
	g.race = map[int]int{Cabbage: Sheep, Sheep: Wolf}
	g.side = true
	g.cap = 1

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1001) GameImage() string {
	var image string
	if g.image != "" {
		image = g.image
		g.image = ""
	}

	return image
}

// GameScene 游戏场景
func (g *Game1001) GameScene() string {
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
func (g *Game1001) GameTips() string {
	return "运送货物过河后可以带回对岸的羊从而避免竞争"
}

// Strategy 攻略
func (g *Game1001) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nv8h6H-tQ8MnrUrOBH5igpu8"
}

// Remind 提醒
func (g *Game1001) Remind() string {
	if len(g.left) == 0 {
		return "您已通关狼羊菜过河，请通过点击菜单或发送指令选择其它游戏继续挑战"
	}

	return "还未通关狼羊菜过河，开动脑筋继续挑战吧，当然您也可以通过点击菜单或发送指令获取提示和攻略\n\n" + g.GameScene()
}
