package logic

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/panshiqu/goddz/wechat"
	"github.com/seefan/gossdb"
)

// Processor 处理器
type Processor struct {
	ssdb    *gossdb.Connectors // SSDB
	players map[string]*Player // 玩家
	mutex   sync.Mutex         // 玩家锁
}

// 实例
var ins *Processor

// SsdbPool 获取连接池
func (p *Processor) SsdbPool() *gossdb.Connectors {
	return p.ssdb
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
	case RefreshTimer:
		wechat.ATIns().Refresh()
	}
}

// OnScan 扫描到来
func (p *Processor) OnScan(user string, message string) {
	switch message {
	// 睢县公园
	case "SuiXian Park":
		log.Println("SuiXian Park")
	}
}

// OnEvent 事件到来
func (p *Processor) OnEvent(user string, message string) {
	// 管理者
	if user == AdminOpenID {
		switch message {
		// 手动刷新
		case "refresh":
			wechat.ATIns().Refresh()
			return

		// 运行状态
		case "status":
			var total int
			var buf bytes.Buffer

			// 加锁
			p.mutex.Lock()

			for _, v := range p.players {
				buf.WriteString(v.GetOpenID())
				buf.WriteString(":")
				buf.WriteString(strconv.Itoa(v.GetCnt()))
				buf.WriteString("\n")
				total += v.GetCnt()
			}

			buf.WriteString("Total:")
			buf.WriteString(strconv.Itoa(total))
			wechat.PushTextMessage(user, buf.String())

			// 解锁
			p.mutex.Unlock()
			return
		}
	}

	// 加锁
	p.mutex.Lock()

	// 查找玩家
	player, ok := p.players[user]
	if !ok {
		// 创建玩家
		player = new(Player)
		player.SetOpenID(user)
		p.players[user] = player
	}

	// 解锁
	p.mutex.Unlock()

	// 通知事件
	player.OnEvent(message)
}

// OnSubscribe 订阅到来
func (p *Processor) OnSubscribe(user string, message string) {
	wechat.PushTextMessage(user, "欢迎关注休闲益智游戏服务号，我们将定期更新休闲益智游戏供你挑战")

	// 未关注扫描带参数二维码
	switch strings.TrimLeft(message, "qrscene_") {
	case "SuiXian Park":
		log.Println("SuiXian Park")
	}
}

// PIns 单例模式
func PIns() *Processor {
	if ins == nil {
		ins = new(Processor)
		ins.players = make(map[string]*Player)
	}

	return ins
}
