package logic

import "strings"

const (
	// Person 人
	Person int = 1

	// Ghost 鬼
	Ghost int = 2
)

// Game1002 游戏
type Game1002 struct {
	left    []int          // 左岸
	right   []int          // 右岸
	carry   []int          // 携带
	voice   map[string]KV  // 语音
	mapping map[int]string // 映射
	name    map[string]int // 名称
	side    bool           // 位置
	cap     int            // 容量
}

// Background 背景
func (g *Game1002) Background() string {
	return `三人三鬼准备过河，人和鬼都会开船，可是岸边只有一条仅能同时承载两个货物的空船，而且鬼比人多的时候鬼将吃人。你能指导他们安全过河吗？`
}

// Description 描述
func (g *Game1002) Description() string {
	return `操作：装(1)、卸(2)、过河(3)
货物：人(1)、鬼(2)

操作和货物之间用点号或空格分隔
装人上船请输入：装。人 或 1.1
卸人下船请输入：卸。人 或 2.1
过河请输入：过河 或 3

支持语音识别，请说普通话
装人上船请发送语音：装人
卸人下船请发送语音：卸人
过河请发送语音：过河`
}

// OnGameEvent 游戏事件
func (g *Game1002) OnGameEvent(event string) string {
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

		lp := Count(left, Person)
		if lp > 0 && Count(left, Ghost) > lp {
			return "左岸的人被鬼吃掉了"
		}

		rp := Count(right, Person)
		if rp > 0 && Count(right, Ghost) > rp {
			return "右岸的人被鬼吃掉了"
		}

		g.left = left
		g.right = right
		g.side = side
		g.carry = []int{}

		if len(g.left) == 0 {
			return "Well Done"
		}
	}

	return g.GameScene()
}

// OnGameStart 游戏开始
func (g *Game1002) OnGameStart() string {
	g.left = []int{Person, Person, Person, Ghost, Ghost, Ghost}
	g.right = []int{}
	g.carry = []int{}
	g.voice = map[string]KV{"guohe": {3, 0}, "zhuangren": {1, 1}, "zhuanggui": {1, 2}, "xieren": {2, 1}, "xiegui": {2, 2}}
	g.mapping = map[int]string{Person: "人", Ghost: "鬼"}
	g.name = map[string]int{"装": Put, "卸": Get, "过河": Go, "人": Person, "鬼": Ghost, "1": 1, "2": 2, "3": 3}
	g.side = true
	g.cap = 2

	return g.GameScene()
}

// GameScene 游戏场景
func (g *Game1002) GameScene() string {
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
	scene += "，请输入"

	return scene
}

// GameTips 提示
func (g *Game1002) GameTips() string {
	return "太简单，没有办法提示"
}
