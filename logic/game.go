package logic

// Game 游戏
type Game interface {
	Description() string

	OnGameEvent(string) string

	OnGameStart() string
}
