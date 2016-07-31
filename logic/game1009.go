package logic

import (
	"strconv"
	"strings"
)

// Game1009 游戏
type Game1009 struct {
	light [][]bool       // 灯阵
	voice map[string]KV  // 语音
	name  map[string]int // 名称
}

// Background 背景
func (g *Game1009) Background() string {
	return `一组3X3灯阵，开关任意一盏灯，都将同时触发其上下左右灯的开关，你能把灯全打开吗？`
}

// Description 描述
func (g *Game1009) Description() string {
	return `无需输入 开 或 关 操作
打开或关闭左上角灯请输入：一。一 或 1.1
打开或关闭右上角灯请输入：三。一 或 3.1

支持语音识别，请说普通话
打开或关闭左上角灯请发送语音：一一
打开或关闭右上角灯请发送语音：三一`
}

// OnGameEvent 游戏事件
func (g *Game1009) OnGameEvent(event string) string {
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

		if len(events) != 2 {
			return "非法操作"
		}

		if cmd, ok = g.name[events[0]]; !ok {
			return "非法X坐标"
		}

		if which, ok = g.name[events[1]]; !ok {
			return "非法Y坐标"
		}
	}

	cmd--
	which--

	// 中
	g.light[which][cmd] = !g.light[which][cmd]

	// 上
	if which-1 >= 0 {
		g.light[which-1][cmd] = !g.light[which-1][cmd]
	}

	// 下
	if which+1 < len(g.light) {
		g.light[which+1][cmd] = !g.light[which+1][cmd]
	}

	// 左
	if cmd-1 >= 0 {
		g.light[which][cmd-1] = !g.light[which][cmd-1]
	}

	// 右
	if cmd+1 < len(g.light[0]) {
		g.light[which][cmd+1] = !g.light[which][cmd+1]
	}

	if g.IsSucceed() {
		return g.GameScene() + "恭喜过关"
	}

	return g.GameScene()
}

// IsSucceed 是否成功
func (g *Game1009) IsSucceed() bool {
	for _, v := range g.light {
		for _, vv := range v {
			if !vv {
				return false
			}
		}
	}

	return true
}

// OnGameStart 游戏开始
func (g *Game1009) OnGameStart() string {
	g.light = [][]bool{{false, false, false}, {false, false, false}, {false, false, false}}
	g.voice = map[string]KV{"yiyi": {1, 1}, "yier": {1, 2}, "yisan": {1, 3}, "eryi": {2, 1}, "erer": {2, 2}, "ersan": {2, 3}, "sanyi": {3, 1}, "saner": {3, 2}, "sansan": {3, 3}}
	g.name = map[string]int{"一": 1, "二": 2, "三": 3, "1": 1, "2": 2, "3": 3}

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1009) GameImage() string {
	return ""
}

// GameScene 游戏场景
func (g *Game1009) GameScene() string {
	scene := "灯阵信息：\n  X1 X2 X3\n"
	for k, v := range g.light {
		for kk, vv := range v {
			if kk == 0 {
				scene += "Y"
				scene += strconv.Itoa(k + 1)
			}

			if vv {
				scene += "亮"
			} else {
				scene += "暗"
			}

			if kk+1 != len(v) {
				scene += " "
			}
		}

		scene += "\n"
	}

	return scene
}

// GameTips 提示
func (g *Game1009) GameTips() string {
	return "太简单，没有办法提示"
}

// Strategy 攻略
func (g *Game1009) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nv2jGS9zF5wSQxfvoPlDpfNc"
}
