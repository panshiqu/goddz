package logic

import (
	"bytes"
	"database/sql"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/panshiqu/goddz/wechat"
	"github.com/seefan/gossdb"
)

// Processor 处理器
type Processor struct {
	db      *sql.DB            // DB
	ssdb    *gossdb.Connectors // SSDB
	players map[string]*Player // 玩家
	mutex   sync.Mutex         // 玩家锁
}

// 实例
var ins *Processor

// GetDB 获取数据库
func (p *Processor) GetDB() *sql.DB {
	return p.db
}

// SsdbPool 获取连接池
func (p *Processor) SsdbPool() *gossdb.Connectors {
	return p.ssdb
}

// InitDB 初始化DB
func (p *Processor) InitDB() bool {
	var err error
	if p.db, err = sql.Open("mysql", DataSourceName); err != nil {
		log.Println("sql.Open ", err)
		return false
	}

	if err = p.db.Ping(); err != nil {
		log.Println("db.Ping ", err)
		return false
	}

	return true
}

// InitSSDB 初始化SSDB
func (p *Processor) InitSSDB() bool {
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
			wechat.PushTextMessage(AdminOpenID, wechat.ATIns().GetAT())
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

		player.Init()
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
	wechat.PushTextMessage(user, WelcomeMessage)
	wechat.PushMpnewsMessage(user, GuideMpnews)

	// 查询数据库
	var status int
	promotion := strings.TrimLeft(message, "qrscene_")
	if err := p.db.QueryRow("SELECT SubscribeStatus FROM UserInfo WHERE OpenID = ?", user).Scan(&status); err == nil {
		// 更新数据
		if _, err := p.db.Exec("UPDATE UserInfo SET PromotionID = ?, SubscribeStatus = TRUE WHERE OpenID = ?", promotion, user); err != nil {
			log.Println("UPDATE", err)
		}
	} else if err == sql.ErrNoRows {
		// 插入数据
		if _, err := p.db.Exec("INSERT INTO UserInfo (OpenID, PromotionID) VALUES (?, ?)", user, promotion); err != nil {
			log.Println("INSERT", err)
		}
	} else {
		log.Println("SELECT", err)
	}
}

// OnUnSubscribe 取消订阅
func (p *Processor) OnUnSubscribe(user string) {
	if _, err := p.db.Exec("UPDATE UserInfo SET SubscribeStatus = FALSE WHERE OpenID = ?", user); err != nil {
		log.Println("UPDATE", err)
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
