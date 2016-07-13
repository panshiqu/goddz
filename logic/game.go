package logic

// Game 游戏
type Game interface {
	Description() string       // 描述
	OnGameStart() string       // 游戏开始
	OnGameEvent(string) string // 游戏事件
}
