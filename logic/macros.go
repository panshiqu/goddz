package logic

const (
	// Put 装
	Put int = 1

	// Get 卸
	Get int = 2

	// Go 过河
	Go int = 3
)

const (
	// RefreshTimer 刷新
	RefreshTimer int64 = 1
)

const (
	// AdminOpenID 管理员
	AdminOpenID string = "o0qWoxE_BrLlGqXE2wJU7SZ01lh0"

	// GuideMpnews 引导图文
	GuideMpnews string = "pEnTAPWdIFaIB0fVJT1nvwjL4n1EzYud7Kzkq8Vjz28"

	// WelcomeMessage 欢迎消息
	WelcomeMessage string = "欢迎关注休闲益智游戏服务号，我们将定期更新休闲益智游戏供你挑战，网页版游戏请访问官网：www.iplaygame.com.cn"

	// DataSourceName 数据源名称
	DataSourceName string = "root:abcwoaini116227@tcp(127.0.0.1:3306)/game?charset=utf8"
)

// KV 简单结构
type KV struct {
	K int
	V int
}
