package logic

import (
	"wechat"
)

// Processor 处理器
type Processor struct{}

// 实例
var ins *Processor

// OnTimer 定时器到期
func (p *Processor) OnTimer(tid int64, param interface{}) {
	switch tid {
	case 1:
		wechat.ATIns().Refresh()
	case 2:
		wechat.PushTextMessage("oilXLwUJYuhN9-ml2aq2yIMZFByo", "定时测试")
	}
}

// PIns 单例模式
func PIns() *Processor {
	if ins == nil {
		ins = new(Processor)
	}

	return ins
}
