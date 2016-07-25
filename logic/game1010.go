package logic

// Game1010 游戏
type Game1010 struct {
	index int      // 编号
	image []string // 图片
	scene []string // 场景
}

// Background 背景
func (g *Game1010) Background() string {
	return `幸苦收集供君欣赏`
}

// Description 描述
func (g *Game1010) Description() string {
	return `回复任意信息获取下一张`
}

// OnGameEvent 游戏事件
func (g *Game1010) OnGameEvent(event string) string {
	g.index++

	if g.index >= len(g.scene) {
		return "已浏览完图集"
	}

	return g.GameScene()
}

// OnGameStart 游戏开始
func (g *Game1010) OnGameStart() string {
	g.image = []string{"pEnTAPWdIFaIB0fVJT1nv_BxXfl8XLVeNrJD8X7akvI",
		"pEnTAPWdIFaIB0fVJT1nvzoRI9e9CkdMqkNtwLkOl1w",
		"pEnTAPWdIFaIB0fVJT1nv_CV2jGbKJfzDrZz6UHAMM4",
		"pEnTAPWdIFaIB0fVJT1nv6Bmf2WCYJd9zR4QvJUvcfw",
		"pEnTAPWdIFaIB0fVJT1nv_Jo9KMbkehLj3K-Y1DT6qw",
		"pEnTAPWdIFaIB0fVJT1nvyVyAb36BEDpRnrPWxYjDUA",
		"pEnTAPWdIFaIB0fVJT1nv1sXMHXLCXZ4o4VBTh2NAhY",
		"pEnTAPWdIFaIB0fVJT1nv4L3ufKUO1Graxh2B8Xnkpw",
		"pEnTAPWdIFaIB0fVJT1nvwJAQbrhdzLnIZ8kqRpnfic",
		"pEnTAPWdIFaIB0fVJT1nv7h6Pop0IvzW5HRAYRB4EDc",
		"pEnTAPWdIFaIB0fVJT1nvzZ1R4M_EttoLzx9vxw9K98",
		"pEnTAPWdIFaIB0fVJT1nv8X_ok9PVewRFrzNeHDT3NI",
		"pEnTAPWdIFaIB0fVJT1nv7hs2Y-YpSIYmJKJ4vk-GX8",
		"pEnTAPWdIFaIB0fVJT1nv6gBSuVyMg_dL4DE9B_GSaw",
		"pEnTAPWdIFaIB0fVJT1nvyE_rRlMohBNTV4IBuv1hZ4",
		"pEnTAPWdIFaIB0fVJT1nv4Ep1tMzQBcArAo30ZRkASQ",
		"pEnTAPWdIFaIB0fVJT1nv-bnFUMvXYdpJQ1pT28-ZgQ",
		"pEnTAPWdIFaIB0fVJT1nv9L2RB6SxAqDK2t8JbmE4sU",
		"pEnTAPWdIFaIB0fVJT1nvy7K-liQDFeuvqigt1v3oQs",
		"pEnTAPWdIFaIB0fVJT1nv_4XbfbdprK-2YzmpcQGRfM",
		"pEnTAPWdIFaIB0fVJT1nv_5Nc_jxSrhfKhOxpP2u72s",
		"pEnTAPWdIFaIB0fVJT1nv-2LRwyRs5czcxd4MG2rOmI",
		"pEnTAPWdIFaIB0fVJT1nvxPkoIKAvW39Mce1EPUEO98",
		"pEnTAPWdIFaIB0fVJT1nv1qHGyVVGh1B2DiWr8tkKKM",
		"pEnTAPWdIFaIB0fVJT1nv2JzE9viZxQPYgu-0zxmLTg",
		"pEnTAPWdIFaIB0fVJT1nv_WTmh2NuuXDNbXI_L34My8",
		"pEnTAPWdIFaIB0fVJT1nv1Hw5Yr9i4TWh5ANVWf5-70",
		"pEnTAPWdIFaIB0fVJT1nv625NzQfnvfNzGdAuPFXZ3w"}

	g.scene = []string{"如果你看到的是顺时针转，说明你用的是右脑。如果你看到的是逆时针转，说明你用的是左脑。据说是耶鲁大学耗时5年的研究成果，有14%的美国人两个方向都可以看到。",
		"不敢相信图中的直线是平行的，不过它就是平行的。",
		"线AB和线CD长度完全相等，虽然它们看起来相差很大。",
		"除了叶子，您还看见了什么？",
		"大象的腿，怎么这么多？",
		"盯着黑点看，发现了什么？",
		"盯着看，前后晃动身体。",
		"动的还是静的？",
		"多少黑点？",
		"更多的点一同出现。",
		"据说看到九张脸的人智商有180。",
		"据说能看见十一张脸的智商有...你想是多少就多少吧。",
		"看不明白。晕了吧！",
		"看起来是一个旋涡，但实际上是由许多同心圆组成的。",
		"流传于网络的千古迷题，无数的人仍在寻找真相…",
		"哪条红线更长？两条红线完全等长。",
		"难道这些点是徘徊在横纵间的灵魂？",
		"你不相信火星木匠来过地球？看看这幅图吧，这是唯一的证据。",
		"前后伸伸头，左右挪挪头，天哪！好美的彩虹！",
		"前后伸伸头，左右挪挪头，天哪！水波动起来！",
		"前后伸伸头，左右挪挪头，天哪！图在动！",
		"曲线正方形？这些是完全的正方形吗？正方形看起来是变形了，但其实它们的边线都是笔直的。",
		"人是贴图里的还是真实的？",
		"狮子在哪里？",
		"十二个还是十三个？",
		"是不是觉得两条黑线向外突。",
		"上下滑动屏幕，灵异现象发生了…",
		"柱子是圆的还是方的？"}

	return g.GameScene()
}

// GameImage 游戏图片
func (g *Game1010) GameImage() string {
	if g.index < len(g.image) {
		return g.image[g.index]
	}

	return ""
}

// GameScene 游戏场景
func (g *Game1010) GameScene() string {
	if g.index < len(g.scene) {
		return g.scene[g.index]
	}

	return ""
}

// GameTips 提示
func (g *Game1010) GameTips() string {
	return "不需要提示"
}
