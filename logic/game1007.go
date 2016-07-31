package logic

import (
	"strconv"
	"strings"
)

const (
	// Bottle1 1
	Bottle1 int = 1

	// Bottle2 2
	Bottle2 int = 2

	// Bottle3 3
	Bottle3 int = 3
)

// Game1007 游戏
type Game1007 struct {
	capacity []int          // 容量
	bottle   []int          // 瓶子
	voice    map[string]KV  // 语音
	name     map[string]int // 名称
}

// Background 背景
func (g *Game1007) Background() string {
	return `哥哥和弟弟打开了一瓶10L的饮料，想着均分后再喝才算公平，在家翻箱倒柜后仅找到3L和7L的标准容器，这下可把哥弟俩难住了，聪明的你能帮帮他们吗？`
}

// Description 描述
func (g *Game1007) Description() string {
	return `瓶子：一(1)、二(2)、三(3)

无需输入 倒 操作
一号倒入二号请输入：一。二 或 1.2

支持语音识别，请说普通话
一号倒入二号请发送语音：一二`
}

// OnGameEvent 游戏事件
func (g *Game1007) OnGameEvent(event string) string {
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
			return "非法瓶子"
		}

		if which, ok = g.name[events[1]]; !ok {
			return "非法瓶子"
		}

		if cmd == which {
			return "瓶子相同"
		}
	}

	if g.bottle[cmd-1] == 0 {
		return "准备倒出的瓶子为空"
	}

	if g.bottle[which-1] >= g.capacity[which-1] {
		return "准备倒入的瓶子已满"
	}

	diff := g.bottle[cmd-1]
	if diff > g.capacity[which-1]-g.bottle[which-1] {
		diff = g.capacity[which-1] - g.bottle[which-1]
	}

	g.bottle[cmd-1] -= diff
	g.bottle[which-1] += diff

	if g.bottle[1] == 5 && g.bottle[2] == 5 {
		return g.GameScene() + "\n恭喜过关"
	}

	return g.GameScene()
}

// OnGameStart 游戏开始
func (g *Game1007) OnGameStart() string {
	g.capacity = []int{3, 7, 10}
	g.bottle = []int{0, 0, 10}
	g.voice = map[string]KV{"yier": {1, 2}, "yisan": {1, 3}, "eryi": {2, 1}, "ersan": {2, 3}, "sanyi": {3, 1}, "saner": {3, 2}}
	g.name = map[string]int{"一": Bottle1, "二": Bottle2, "三": Bottle3, "1": 1, "2": 2, "3": 3}

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1007) GameImage() string {
	return ""
}

// GameScene 游戏场景
func (g *Game1007) GameScene() string {
	scene := "瓶子信息："
	for k, v := range g.bottle {
		scene += strconv.Itoa(v)
		scene += "/"
		scene += strconv.Itoa(g.capacity[k])

		if k+1 != len(g.bottle) {
			scene += "、"
		}
	}

	return scene
}

// GameTips 提示
func (g *Game1007) GameTips() string {
	return "太简单，没有办法提示"
}

// Strategy 攻略
func (g *Game1007) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nv0Ie66Y8jxaNAwxGWtw_svs"
}
