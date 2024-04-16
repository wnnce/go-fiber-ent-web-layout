package limiter

import (
	"context"
	"fmt"
	"path"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tokenBucket := &TokenBucket{
		name:       "demo-token",
		maxNum:     1000,
		avail:      20,
		interval:   1 * time.Second,
		releaseNum: 50,
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	tokenBucket.TimingReleaseToken(ctx)
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Printf("limit: %v\n", tokenBucket.DoLimit("demo"))
			case <-ctx.Done():
				return
			}
		}
	}()
	for _ = range ctx.Done() {
		break
	}
}

func TestSlidingWindow(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	window := NewSlidingWindow(SlidingConfig{
		Interval:  3 * time.Second,
		WindowNum: 5,
	})
	window.TimingSideWindow(ctx)
	go func() {
		ticker := time.NewTicker(30 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Printf("demo: %v\n", window.DoLimit("demo"))
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(80 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Printf("test: %v\n", window.DoLimit("test"))
			case <-ctx.Done():
				return
			}
		}
	}()
	time.Sleep(30 * time.Second)
	cancel()
}

func TestPathMatch(t *testing.T) {
	result, _ := path.Match("/demo", "/demo")
	fmt.Println(result)
	result, _ = path.Match("/demo", "/demo/1")
	fmt.Println(result)
	result, _ = path.Match("demo", "/demo/1")
	fmt.Println(result)
	result, _ = path.Match("/demo/*", "/demo/1")
	fmt.Println(result)
	result, _ = path.Match("/demo/*/cc", "/demo/asda/cc")
	fmt.Println(result)
	result, _ = path.Match("/demo/*/*/asda", "/demo/asda/cc/asda")
	fmt.Println(result)
	result, _ = path.Match("/demo/*", "/demo/asda/cc/asda")
	fmt.Println(result)
	result, _ = path.Match("/demo/*/?cc", "/demo/asda/acc")
	fmt.Println(result)
}

func TestIPv4Match_DoMatch(t *testing.T) {
	match := NewIPMatch()
	fmt.Println(match.Match("127.0.0.1", "127.0.0.1"))
	fmt.Println(match.Match("127.0.*.1", "127.0.0.1"))
	fmt.Println(match.Match("127.0.*.1", "127.0.0"))
}
