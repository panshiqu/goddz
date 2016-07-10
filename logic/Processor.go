package logic

import (
	"log"
	"math/rand"
	"time"

	"github.com/panshiqu/goddz/wechat"
	"github.com/seefan/gossdb"
)

// Processor 处理器
type Processor struct {
	players map[string]*Player
	ssdb    *gossdb.Connectors
	random  *rand.Rand
}

// 实例
var ins *Processor

// SsdbPool 获取连接池
func (p *Processor) SsdbPool() *gossdb.Connectors {
	return p.ssdb
}

// Random 获取随机数
func (p *Processor) Random() *rand.Rand {
	return p.random
}

// Init 初始化
func (p *Processor) Init() bool {
	var err error
	p.ssdb, err = gossdb.NewPool(&gossdb.Config{
		Host:             "127.0.0.1",
		Port:             8888,
		MaxPoolSize:      50,
		MinPoolSize:      5,
		AcquireIncrement: 5,
	})

	if err != nil {
		log.Println("gossdb.NewPool ", err)
		return false
	}

	return true
}

// OnTimer 定时器到期
func (p *Processor) OnTimer(tid int64, param interface{}) {
	switch tid {
	case 1:
		wechat.ATIns().Refresh()
	case 2:
		wechat.PushTextMessage("oilXLwUJYuhN9-ml2aq2yIMZFByo", "定时测试")
	}
}

// OnEvent 事件到来
func (p *Processor) OnEvent(user string, message string) {
	// 自定义菜单
	if message == "can ju" {
		wechat.PushTextMessage(user, "敬请期待")
		return
	} else if message == "contact us" {
		wechat.PushTextMessage(user, "联系电话：13526535277")
		return
	}

	// 查找用户
	player, ok := p.players[user]
	if !ok {
		// 创建用户
		player = new(Player)
		player.SetOpenID(user)
		p.players[user] = player
	}

	// 通知事件
	player.OnEvent(message)
}

// PIns 单例模式
func PIns() *Processor {
	if ins == nil {
		ins = new(Processor)
		ins.players = make(map[string]*Player)
		ins.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return ins
}
