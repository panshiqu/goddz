package logic

import "strconv"

// Game1008 游戏
type Game1008 struct {
	gold  []int          // 金条
	name  map[string]int // 名称
	image string         // 图片
}

// GetID 获取编号
func (g *Game1008) GetID() int {
	return 1008
}

// IsSucceed 是否成功
func (g *Game1008) IsSucceed() bool {
	if Contain(g.gold, 1) && Contain(g.gold, 2) && Contain(g.gold, 4) {
		return true
	}

	return false
}

// Background 背景
func (g *Game1008) Background() string {
	return `一名短工打理院落，需工作七天，土豪仅从自己金库里拿出一根金条用于发工资，短工日工资正好是1/7金条，假若土豪能准确切出他想要的比例，他能只切两刀就能发放这七天的工资吗？`
}

// Description 描述
func (g *Game1008) Description() string {
	return `请发送数字、数字中文、数字语音切金条，请始终保持待切金条在左侧，若发送 3（三、语音） 将会把金条切成 3|4 两份，若再输入 1（一、语音） 将会把金条切成 1|2|4 三份`
}

// OnGameEvent 游戏事件
func (g *Game1008) OnGameEvent(event string) string {
	if len(g.gold) >= 3 {
		return "已切够两刀"
	}

	which, ok := g.name[event]
	if !ok {
		return "不能这样切"
	}

	if g.gold[0] <= which {
		return "金条不够切"
	}

	tmp := []int{which, g.gold[0] - which}
	tmp = append(tmp, g.gold[1:]...)
	g.gold = tmp

	if Contain(g.gold, 1) && Contain(g.gold, 2) && Contain(g.gold, 4) {
		return g.GameScene() + "\n\n恭喜过关，但是你真的知道怎么发工资了吗？"
	}

	return g.GameScene()
}

// OnGameStart 游戏开始
func (g *Game1008) OnGameStart() string {
	g.gold = []int{7}
	g.name = map[string]int{"yi": 1, "er": 2, "san": 3, "si": 4, "wu": 5, "liu": 6, "一": 1, "二": 2,
		"三": 3, "四": 4, "五": 5, "六": 6, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6}
	g.image = "pEnTAPWdIFaIB0fVJT1nvzZY79hH4vR1rDGhDtrvtD0"

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1008) GameImage() string {
	var image string
	if g.image != "" {
		image = g.image
		g.image = ""
	}

	return image
}

// GameScene 游戏场景
func (g *Game1008) GameScene() string {
	scene := "金条信息："
	for k, v := range g.gold {
		scene += strconv.Itoa(v)
		if k+1 != len(g.gold) {
			scene += "|"
		}
	}

	return scene
}

// GameTips 提示
func (g *Game1008) GameTips() string {
	return "操作说明中好像有答案哦"
}

// Strategy 攻略
func (g *Game1008) Strategy() string {
	return "pEnTAPWdIFaIB0fVJT1nv126xbwMXaiLDqG1bkCr8xQ"
}

// Remind 提醒
func (g *Game1008) Remind() string {
	if Contain(g.gold, 1) && Contain(g.gold, 2) && Contain(g.gold, 4) {
		return "您已通关切分金条发工资，请通过点击菜单或发送指令选择其它游戏继续挑战"
	}

	return "还未通关切分金条发工资，开动脑筋继续挑战吧，当然您也可以通过点击菜单或发送指令获取提示和攻略\n\n" + g.GameScene()
}
