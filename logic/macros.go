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
)

// KV 简单结构
type KV struct {
	K int
	V int
}
