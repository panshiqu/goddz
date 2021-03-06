package logic

// Game 游戏
type Game interface {
	GetID() int                // 编号
	IsSucceed() bool           // 成功
	Background() string        // 背景
	Description() string       // 描述
	OnGameStart() string       // 游戏开始
	OnGameEvent(string) string // 游戏事件
	GameImage() string         // 游戏图片
	GameScene() string         // 场景
	GameTips() string          // 提示
	Strategy() string          // 攻略
	Remind() string            // 提醒
}
