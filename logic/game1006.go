package logic

import "strconv"

const (
	// Frog1 1
	Frog1 int = 1

	// Frog2 2
	Frog2 int = 2

	// Frog3 3
	Frog3 int = 3

	// Frog4 4
	Frog4 int = 4

	// Frog5 5
	Frog5 int = 5

	// Frog6 6
	Frog6 int = 6
)

// Game1006 游戏
type Game1006 struct {
	frog  []int          // 青蛙
	name  map[string]int // 名称
	image string         // 图片
}

// GetID 获取编号
func (g *Game1006) GetID() int {
	return 1006
}

// IsSucceed 是否成功
func (g *Game1006) IsSucceed() bool {
	if g.frog[0] == 4 && g.frog[1] == 5 && g.frog[2] == 6 &&
		g.frog[4] == 1 && g.frog[5] == 2 && g.frog[6] == 3 {
		return true
	}

	return false
}

// Background 背景
func (g *Game1006) Background() string {
	return `两队青蛙狭路相逢，(1.2.3)向右移动，(4.5.6)向左移动，移动的规则就是跳棋游戏的规则，你能指挥它们顺利跳过吗？`
}

// Description 描述
func (g *Game1006) Description() string {
	return `青蛙：一(1)、二(2)、三(3)、四(4)、五(5)、六(6)

无需输入 跳 操作
跳一号青蛙请输入：一 或 1

支持语音识别，请说普通话
跳一号青蛙请发送语音：一`
}

// OnGameEvent 游戏事件
func (g *Game1006) OnGameEvent(event string) string {
	which, ok := g.name[event]
	if !ok {
		return "非法青蛙"
	}

	pos := Index(g.frog, which)
	if pos == -1 {
		return "出现严重问题"
	}

	if which <= Frog3 {
		if pos+1 < len(g.frog) && g.frog[pos+1] == 0 {
			g.frog[pos], g.frog[pos+1] = g.frog[pos+1], g.frog[pos]
		} else if pos+2 < len(g.frog) && g.frog[pos+2] == 0 {
			g.frog[pos], g.frog[pos+2] = g.frog[pos+2], g.frog[pos]
		} else {
			return "无法向右跳动青蛙"
		}
	} else if which >= Frog4 {
		if pos-1 >= 0 && g.frog[pos-1] == 0 {
			g.frog[pos], g.frog[pos-1] = g.frog[pos-1], g.frog[pos]
		} else if pos-2 >= 0 && g.frog[pos-2] == 0 {
			g.frog[pos], g.frog[pos-2] = g.frog[pos-2], g.frog[pos]
		} else {
			return "无法向左跳动青蛙"
		}
	}

	if g.frog[0] == 4 && g.frog[1] == 5 && g.frog[2] == 6 &&
		g.frog[4] == 1 && g.frog[5] == 2 && g.frog[6] == 3 {
		return g.GameScene() + "\n\n恭喜过关"
	}

	return g.GameScene()
}

// OnGameStart 游戏开始
func (g *Game1006) OnGameStart() string {
	g.frog = []int{1, 2, 3, 0, 4, 5, 6}
	g.name = map[string]int{"yi": 1, "er": 2, "san": 3, "si": 4, "wu": 5, "liu": 6, "一": Frog1, "二": Frog2,
		"三": Frog3, "四": Frog4, "五": Frog5, "六": Frog6, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6}
	g.image = "pEnTAPWdIFaIB0fVJT1nvxQDDKYATgNCzGUYo6OwbYg"

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1006) GameImage() string {
	var image string
	if g.image != "" {
		image = g.image
		g.image = ""
	}

	return image
}

// GameScene 游戏场景
func (g *Game1006) GameScene() string {
	scene := "青蛙信息："
	for k, v := range g.frog {
		if v != 0 {
			scene += strconv.Itoa(v)
		} else {
			scene += "_"
		}

		if k+1 != len(g.frog) {
			scene += " "
		}
	}

	return scene
}

// GameTips 提示
func (g *Game1006) GameTips() string {
	return "跳之前先把对面的青蛙跳到自己面前"
}

// Strategy 攻略
func (g *Game1006) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nv30aGZLKjR5pp2KEonUk_Ck"
}

// Remind 提醒
func (g *Game1006) Remind() string {
	if g.frog[0] == 4 && g.frog[1] == 5 && g.frog[2] == 6 &&
		g.frog[4] == 1 && g.frog[5] == 2 && g.frog[6] == 3 {
		return "您已通关青蛙直线跳棋，请通过点击菜单或发送指令选择其它游戏继续挑战"
	}

	return "还未通关青蛙直线跳棋，开动脑筋继续挑战吧，当然您也可以通过点击菜单或发送指令获取提示和攻略\n\n" + g.GameScene()
}
