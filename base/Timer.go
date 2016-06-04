package base

import (
	"log"
	"logic"
	"sync"
	"time"
)

// TimerElement 定时器
type TimerElement struct {
	tid      int64         // 编号
	param    interface{}   // 参数
	begin    time.Time     // 开始时间
	timer    *time.Timer   // 定时实例
	duration time.Duration // 持续时间
	repeated bool          // 是否重复
}

// TimerManager 定时器管理
type TimerManager struct {
	tm    map[int64]*TimerElement // 定时器管理
	mutex sync.Mutex              // 互斥锁
}

// 实例
var ins *TimerManager

// AddTimer 添加定时器
func (t *TimerManager) AddTimer(tid int64, duration time.Duration, repeated bool, param interface{}) {
	// 加锁
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 是否存在
	te, ok := t.tm[tid]
	if !ok {
		// 创建定时器
		te = new(TimerElement)

		// 设置参数
		te.tid = tid
		te.param = param
		te.duration = duration
		te.repeated = repeated
		te.timer = time.AfterFunc(duration, func() {
			t.OnTimer(tid, param)
		})

		// 增加定时器管理
		t.tm[tid] = te
	}

	// 不管怎样都做
	te.begin = time.Now()
	te.timer.Reset(te.duration)
}

// DelTimer 删除定时器
func (t *TimerManager) DelTimer(tid int64) {
	// 加锁
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 是否存在
	te, ok := t.tm[tid]
	if ok {
		// 停止定时器
		te.timer.Stop()

		// 删除定时器
		delete(t.tm, tid)
	}
}

// IsExist 是否存在
func (t *TimerManager) IsExist(tid int64) bool {
	// 加锁
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 是否存在
	_, ok := t.tm[tid]

	return ok
}

// GetSurplus 获取剩余时间
func (t *TimerManager) GetSurplus(tid int64) time.Duration {
	// 加锁
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 是否存在
	te, ok := t.tm[tid]
	if ok {
		// 返回定时器剩余时间
		return te.begin.Add(te.duration).Sub(time.Now())
	}

	return time.Duration(0)
}

// resetTimer 重置定时器
func (t *TimerManager) resetTimer(tid int64) {
	// 加锁
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 是否存在
	te, ok := t.tm[tid]
	if ok {
		// 定时器重复
		if te.repeated {
			// 更新开始时间
			te.begin = time.Now()

			// 重置定时器
			te.timer.Reset(te.duration)
		} else {
			// 删除定时器
			delete(t.tm, tid)
		}
	}
}

// OnTimer 定时器到期
func (t *TimerManager) OnTimer(tid int64, param interface{}) {
	// 通知业务逻辑
	log.Println("OnTimer ", tid, param)
	logic.PIns().OnTimer(tid, param)

	// 重置定时器
	t.resetTimer(tid)
}

// TMIns 单例模式
func TMIns() *TimerManager {
	if ins == nil {
		ins = new(TimerManager)
		ins.tm = make(map[int64]*TimerElement)
	}

	return ins
}
