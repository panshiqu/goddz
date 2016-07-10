package logic

import "strings"

// Game1001 游戏
type Game1001 struct {
	left    []int          // 左岸
	right   []int          // 右岸
	carry   int            // 携带
	where   bool           // 船在那里
	mapping map[int]string // 映射
	name    map[string]int // 名称
	race    map[int]int    // 竞争
}

// Description 描述
func (g *Game1001) Description() string {
	return "把卷心菜、小羊、狼运到对岸\n注意人不在的时候小羊会吃掉卷心菜、狼会吃掉小羊\n名词：卷心菜(0)、小羊(1)、狼(2)\n操作：装船(3)、过河(4)\n操作名词之间逗号分隔"
}

// OnGameEvent 游戏事件
func (g *Game1001) OnGameEvent(event string) string {
	var events []string
	if strings.Contains(event, ",") {
		events = strings.Split(event, ",")
	} else if strings.Contains(event, "，") {
		events = strings.Split(event, "，")
	} else {
		events = []string{event}
	}

	cmd, ok := g.name[events[0]]
	if !ok || cmd > 4 || cmd < 3 {
		return "非法指令"
	}

	which := -1
	if len(events) > 1 && cmd == 3 {
		if which, ok = g.name[events[1]]; !ok {
			return "非法货物"
		}
	}

	switch cmd {
	case 3:
		if g.where && !Contain(g.left, which) || !g.where && !Contain(g.right, which) {
			return "请重新装船"
		}

		g.carry = which
	case 4:
		left := g.left
		right := g.right
		where := !g.where

		if where {
			if pos := Index(g.right, g.carry); pos != -1 {
				right = append(g.right[0:pos], g.right[pos+1:]...)
				left = append(g.left, g.carry)
			}

			if ok, k, v := g.RaceDetect(right); !ok {
				return "你右岸的" + g.mapping[k] + "被" + g.mapping[v] + "吃掉了"
			}
		} else {
			if pos := Index(g.left, g.carry); pos != -1 {
				left = append(g.left[0:pos], g.left[pos+1:]...)
				right = append(g.right, g.carry)
			}

			if ok, k, v := g.RaceDetect(left); !ok {
				return "你左岸的" + g.mapping[k] + "被" + g.mapping[v] + "吃掉了"
			}
		}

		g.left = left
		g.right = right
		g.where = where
		g.carry = -1

		if len(g.left) == 0 {
			return g.OnGameOver()
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
	g.left = []int{0, 1, 2}
	g.right = []int{}
	g.carry = -1
	g.where = true
	g.mapping = map[int]string{0: "卷心菜", 1: "小羊", 2: "狼"}
	g.name = map[string]int{"装船": 3, "过河": 4, "卷心菜": 0, "小羊": 1, "狼": 2, "0": 0, "1": 1, "2": 2, "3": 3, "4": 4}
	g.race = map[int]int{0: 1, 1: 2}

	return g.GameScene()
}

// OnGameOver 游戏结束
func (g *Game1001) OnGameOver() string {
	return "Well Done"
}

// GameScene 游戏场景
func (g *Game1001) GameScene() string {
	scene := "左岸："
	for k, v := range g.left {
		if v == g.carry {
			continue
		}

		scene += g.mapping[v]
		if k+1 != len(g.left) {
			scene += "、"
		}
	}

	scene += "\n右岸："
	for k, v := range g.right {
		if v == g.carry {
			continue
		}

		scene += g.mapping[v]
		if k+1 != len(g.right) {
			scene += "、"
		}
	}

	scene += "\n船上：" + g.mapping[g.carry]

	scene += "\n船在" + g.WhereBoat() + "，请操作"

	return scene
}

// WhereBoat 船在那里
func (g *Game1001) WhereBoat() string {
	if g.where {
		return "左岸"
	}

	return "右岸"
}
