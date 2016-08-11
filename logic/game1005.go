package logic

import (
	"strconv"
	"strings"
)

const (
	// Select 选择
	Select int = 1

	// Cancel 取消
	Cancel int = 2

	// Up 上
	Up int = 3

	// Down 下
	Down int = 4
)

const (
	// Elevator1 1
	Elevator1 int = 1

	// Elevator2 2
	Elevator2 int = 2

	// Elevator3 3
	Elevator3 int = 3

	// Elevator4 4
	Elevator4 int = 4

	// Elevator5 5
	Elevator5 int = 5
)

// Game1005 游戏
type Game1005 struct {
	elevator []int          // 电梯
	selected []int          // 已选
	image    string         // 图片
	voice    map[string]KV  // 语音
	name     map[string]int // 名称
	cap      int            // 容量
}

// GetID 获取编号
func (g *Game1005) GetID() int {
	return 1005
}

// Background 背景
func (g *Game1005) Background() string {
	return `一天某小区的五部电梯集体故障，分别停在17、26、20、19、31层，每部电梯里面都困有居民，只有当五部电梯同时停在21-25层时，才能打开电梯门救出被困居民，此时只能任选两部电梯上8层或下13层，该小区最高只到49层（顶层）、最低只到0层（B1层），你能紧急指挥救出那些被困的居民吗？`
}

// Description 描述
func (g *Game1005) Description() string {
	return `操作：选择(1)、取消(2)、上(3)、下(4)
电梯：一(1)、二(2)、三(3)、四(4)、五(5)

操作和电梯之间用点号或空格分隔
选择第一个电梯请输入：选择。一 或 1.1
取消第一个电梯请输入：取消。一 或 2.1
两部电梯上8层请输入：上 或 3
两部电梯下13层请输入：下 或 4

支持语音识别，请说普通话
选择第一个电梯请发送语音：选择一
取消第一个电梯请发送语音：取消一
上请发送语音：上
下请发送语音：下`
}

// OnGameEvent 游戏事件
func (g *Game1005) OnGameEvent(event string) string {
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
		if !ok || cmd < Select || cmd > Down {
			return "非法操作"
		}

		if len(events) > 1 && cmd != Up && cmd != Down {
			if which, ok = g.name[events[1]]; !ok {
				return "非法电梯"
			}
		}
	}

	switch cmd {
	case Select:
		if len(g.selected) >= g.cap {
			return "只能选择两部电梯"
		}

		g.selected = append(g.selected, which)

	case Cancel:
		if len(g.selected) <= 0 {
			return "你都没有选择电梯"
		}

		pos := Index(g.selected, which)
		if pos == -1 {
			return "该部电梯未被选中"
		}

		g.selected = append(g.selected[0:pos], g.selected[pos+1:]...)

	case Up:
		if len(g.selected) != g.cap {
			return "两部电梯同时选中才能运行"
		}

		if g.elevator[g.selected[0]-1]+8 > 49 || g.elevator[g.selected[1]-1]+8 > 49 {
			return "大楼最高只到49层（顶层）"
		}

		g.elevator[g.selected[0]-1] += 8
		g.elevator[g.selected[1]-1] += 8
		g.selected = []int{}

		if g.IsSucceed() {
			return g.GameScene() + "\n\n恭喜过关"
		}

	case Down:
		if len(g.selected) != g.cap {
			return "两部电梯同时选中才能运行"
		}

		if g.elevator[g.selected[0]-1]-13 < 0 || g.elevator[g.selected[1]-1]-13 < 0 {
			return "大楼最低只到0层（B1层）"
		}

		g.elevator[g.selected[0]-1] -= 13
		g.elevator[g.selected[1]-1] -= 13
		g.selected = []int{}

		if g.IsSucceed() {
			return g.GameScene() + "\n\n恭喜过关"
		}
	}

	return g.GameScene()
}

// IsSucceed 是否成功
func (g *Game1005) IsSucceed() bool {
	for _, v := range g.elevator {
		if v < 21 || v > 25 {
			return false
		}
	}

	return true
}

// OnGameStart 游戏开始
func (g *Game1005) OnGameStart() string {
	g.elevator = []int{17, 26, 20, 19, 31}
	g.selected = []int{}
	g.image = "pEnTAPWdIFaIB0fVJT1nv2FEDqqVW1pqtetNOHo4_5k"
	g.voice = map[string]KV{"shang": {3, 0}, "xia": {4, 0}, "xuanzeyi": {1, 1}, "xuanzeer": {1, 2}, "xuanzesan": {1, 3}, "xuanzesi": {1, 4},
		"xuanzewu": {1, 5}, "quxiaoyi": {2, 1}, "quxiaoer": {2, 2}, "quxiaosan": {2, 3}, "quxiaosi": {2, 4}, "quxiaowu": {2, 5}}
	g.name = map[string]int{"选择": Select, "取消": Cancel, "上": Up, "下": Down, "一": Elevator1, "二": Elevator2,
		"三": Elevator3, "四": Elevator4, "五": Elevator5, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5}
	g.cap = 2

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1005) GameImage() string {
	var image string
	if g.image != "" {
		image = g.image
		g.image = ""
	}

	return image
}

// GameScene 游戏场景
func (g *Game1005) GameScene() string {
	scene := "电梯信息："
	for k, v := range g.elevator {
		scene += strconv.Itoa(k + 1)
		scene += "-"
		scene += strconv.Itoa(v)
		if k+1 != len(g.elevator) {
			scene += "、"
		}
	}

	scene += "\n已选电梯："
	for k, v := range g.selected {
		scene += strconv.Itoa(v)
		if k+1 != len(g.selected) {
			scene += "、"
		}
	}

	return scene
}

// GameTips 提示
func (g *Game1005) GameTips() string {
	return "8+8－13=3  8+8-13-13=-10"
}

// Strategy 攻略
func (g *Game1005) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nvyi_VHoQ-1iAB9NOxmtYZ9M"
}

// Remind 提醒
func (g *Game1005) Remind() string {
	if g.IsSucceed() {
		return "您已通关指挥电梯上下，请通过点击菜单或发送指令选择其它游戏继续挑战"
	}

	return "还未通关指挥电梯上下，开动脑筋继续挑战吧，当然您也可以通过点击菜单或发送指令获取提示和攻略\n\n" + g.GameScene()
}
