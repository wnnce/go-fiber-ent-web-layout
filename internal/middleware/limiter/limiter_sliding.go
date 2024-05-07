package limiter

import (
	"container/ring"
	"context"
	"sync"
	"time"
)

type Window struct {
	sync.Map
}

// Count 统计当前窗口中该key的数据
func (b *Window) Count(key string) int {
	if value, ok := b.Load(key); ok {
		return value.(int)
	}
	return 0
}

// Add 将key添加到窗口 不存在就添加 存在就加1
func (b *Window) Add(key string) {
	var newValue int
	if value, ok := b.Load(key); ok {
		newValue = value.(int) + 1
	} else {
		newValue = 1
	}
	b.Store(key, newValue)
}

// SlidingWindow 基于ring环形列表的滑动时间窗口
// 用于统计一段时间内客户端的请求次数以达到限流的目的
type SlidingWindow struct {
	name      string        // 名称
	length    int           // 时间窗口的数量
	interval  time.Duration // 每次时间窗口滑动的间隔
	queue     *ring.Ring    // 环形列表
	mutex     sync.RWMutex  // 读写锁
	threshold int64         // 限流阈值
}

// NewSlidingWindow 创建滑动时间窗口
// length 时间窗口内桶的数量 数量越多统计的请求时间段就越多
func NewSlidingWindow(config SlidingConfig) *SlidingWindow {
	config = slidingConfigDefault(config)
	window := &SlidingWindow{
		name:      config.Name,
		length:    config.WindowNum,
		interval:  config.Interval,
		threshold: config.Threshold,
	}
	queue := ring.New(config.WindowNum)
	for i := 0; i < config.WindowNum; i++ {
		queue.Value = new(Window)
		queue = queue.Next()
	}
	window.queue = queue
	return window
}

// 记录请求
func (w *SlidingWindow) record(key string) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	bucket, _ := w.queue.Value.(*Window)
	bucket.Add(key)
}

// 获取当前所有窗口内该key的请求次数
func (w *SlidingWindow) stats(key string) int64 {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	var sum int64
	w.queue.Do(func(i any) {
		bucket, _ := i.(*Window)
		sum += int64(bucket.Count(key))
	})
	return sum
}

// TimingSideWindow 定时更新时间窗口，使用 interval 参数做为间隔时间
// ctx 监听Done事件 停止定时器并退出任务
func (w *SlidingWindow) TimingSideWindow(ctx context.Context) {
	go func(ctx context.Context) {
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				w.mutex.Lock()
				w.queue = w.queue.Next()
				w.queue.Value = new(Window)
				w.mutex.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
}

func (w *SlidingWindow) DoLimit(key string) bool {
	stats := w.stats(key)
	if stats < w.threshold {
		w.record(key)
		return false
	}
	return true
}

func (w *SlidingWindow) Name() string {
	return w.name
}
