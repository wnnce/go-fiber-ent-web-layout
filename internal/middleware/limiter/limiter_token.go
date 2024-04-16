package limiter

import (
	"context"
	"math"
	"sync/atomic"
	"time"
)

// TokenBucket 令牌桶
type TokenBucket struct {
	name       string        // 自定义名称
	maxNum     int64         // 令牌桶中的最大令牌数量
	avail      int64         // 当前可用的令牌数量
	interval   time.Duration // 释放令牌的间隔时间
	releaseNum int           // 每次释放的令牌数量
}

func NewTokenBucket(config TokenBucketConfig) *TokenBucket {
	config = tokenBucketConfigDefault(config)
	tokenBucket := &TokenBucket{
		name:       config.Name,
		maxNum:     config.MaxNum,
		avail:      config.DefaultAvail,
		interval:   config.ReleaseInterval,
		releaseNum: config.ReleaseNum,
	}
	return tokenBucket
}

// TimingReleaseToken 定时释放令牌
func (t *TokenBucket) TimingReleaseToken(ctx context.Context) {
	go func(ctx context.Context) {
		ticker := time.NewTicker(t.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				sub := t.maxNum - atomic.LoadInt64(&t.avail)
				if sub > 0 {
					addTokenNum := int64(math.Min(float64(sub), float64(t.releaseNum)))
					atomic.AddInt64(&t.avail, addTokenNum)
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
}

func (t *TokenBucket) DoLimit(key string) bool {
	if atomic.LoadInt64(&t.avail) <= 0 {
		return false
	}
	atomic.StoreInt64(&t.avail, t.avail-1)
	return true
}

func (t *TokenBucket) Name() string {
	return t.name
}
