package logic

import (
	"strconv"
	"strings"
)

const (
	// Person1 人
	Person1 int = 1

	// Person3 人
	Person3 int = 2

	// Person6 人
	Person6 int = 3

	// Person8 人
	Person8 int = 4

	// Person12 人
	Person12 int = 5

	// TotalTime 总时
	TotalTime int = 30
)

// Game1004 游戏
type Game1004 struct {
	left    []int          // 左岸
	right   []int          // 右岸
	carry   []int          // 携带
	voice   map[string]KV  // 语音
	mapping map[int]string // 映射
	name    map[string]int // 名称
	need    map[int]int    // 需要
	side    bool           // 位置
	cap     int            // 容量
	use     int            // 已用
}

// Background 背景
func (g *Game1004) Background() string {
	return `漆黑的夜里，一家人需要通过一座独木桥，但独木桥最多承载两人的重量，一家人只有一盏灯，但这盏灯只能使用30秒，每个人过桥所需的时间不同，哥哥1秒、弟弟3秒、妈妈6秒、爸爸8秒、爷爷12秒，如何在灯熄灭前顺利通过独木桥，你能指点一下他们吗？`
}

// Description 描述
func (g *Game1004) Description() string {
	return `操作：装(1)、卸(2)、过桥(3)
家人：哥哥(1)、弟弟(2)、妈妈(3)、爸爸(4)、爷爷(5)
家人支持简写：哥(1)、弟(2)、妈(3)、爸(4)、爷(5)

操作和家人之间用点号或空格分隔
装哥哥上桥请输入：装。哥哥 或 1.1
卸哥哥下桥请输入：卸。哥哥 或 2.1
过桥请输入：过桥 或 3

支持语音识别，请说普通话
装哥哥上船请发送语音：装哥哥 或 装哥
卸哥哥下船请发送语音：卸哥哥 或 卸哥
过桥请发送语音：过桥`
}

// OnGameEvent 游戏事件
func (g *Game1004) OnGameEvent(event string) string {
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
				return "非法家人"
			}
		}
	}

	switch cmd {
	case Put:
		if len(g.carry) >= g.cap {
			return "桥已超载"
		}

		if g.side {
			pos := Index(g.left, which)
			if pos == -1 {
				return "请装左岸家人"
			}

			g.left = append(g.left[0:pos], g.left[pos+1:]...)
		} else {
			pos := Index(g.right, which)
			if pos == -1 {
				return "请装右岸家人"
			}

			g.right = append(g.right[0:pos], g.right[pos+1:]...)
		}

		g.carry = append(g.carry, which)
	case Get:
		if len(g.carry) <= 0 {
			return "桥已为空"
		}

		pos := Index(g.carry, which)
		if pos == -1 {
			return "请卸桥上家人"
		}

		if g.side {
			g.left = append(g.left, which)
		} else {
			g.right = append(g.right, which)
		}

		g.carry = append(g.carry[0:pos], g.carry[pos+1:]...)
	case Go:
		if len(g.carry) <= 0 {
			return "无人掌灯"
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

		var max int
		for _, v := range g.carry {
			need := g.need[v]
			if need > max {
				max = need
			}
		}

		if g.use+max > TotalTime {
			return "You Will Lose"
		}

		g.use += max
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
func (g *Game1004) OnGameStart() string {
	g.left = []int{Person1, Person3, Person6, Person8, Person12}
	g.right = []int{}
	g.carry = []int{}
	g.voice = map[string]KV{"过桥": {3, 0}, "装哥哥": {1, 1}, "装弟弟": {1, 2}, "装妈妈": {1, 3}, "装爸爸": {1, 4}, "装爷爷": {1, 5},
		"装哥": {1, 1}, "装弟": {1, 2}, "装妈": {1, 3}, "装爸": {1, 4}, "装爷": {1, 5},
		"卸哥哥": {2, 1}, "卸弟弟": {2, 2}, "卸妈妈": {2, 3}, "卸爸爸": {2, 4}, "卸爷爷": {2, 5},
		"卸哥": {2, 1}, "卸弟": {2, 2}, "卸妈": {2, 3}, "卸爸": {2, 4}, "卸爷": {2, 5}}
	g.mapping = map[int]string{Person1: "哥哥", Person3: "弟弟", Person6: "妈妈", Person8: "爸爸", Person12: "爷爷"}
	g.name = map[string]int{"装": Put, "卸": Get, "过桥": Go, "哥哥": Person1, "弟弟": Person3, "妈妈": Person6, "爸爸": Person8, "爷爷": Person12,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "哥": Person1, "弟": Person3, "妈": Person6, "爸": Person8, "爷": Person12}
	g.need = map[int]int{Person1: 1, Person3: 3, Person6: 6, Person8: 8, Person12: 12}
	g.side = true
	g.cap = 2

	return g.GameScene()
}

// GameScene 游戏场景
func (g *Game1004) GameScene() string {
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

	scene += "\n桥上："
	for k, v := range g.carry {
		scene += g.mapping[v]
		if k+1 != len(g.carry) {
			scene += "、"
		}
	}

	scene += "\n灯在"
	if g.side {
		scene += "左岸"
	} else {
		scene += "右岸"
	}
	scene += "，剩余 "
	scene += strconv.Itoa(TotalTime - g.use)
	scene += " 秒，请输入"

	return scene
}

// GameTips 提示
func (g *Game1004) GameTips() string {
	return "太简单，没有办法提示"
}
