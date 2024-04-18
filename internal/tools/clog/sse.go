package clog

import (
	"fmt"
	"sync"
	"time"
)

type SSEWriter struct {
	chanCache map[string]chan []byte
	mu        *sync.Mutex
}

var sseWriter *SSEWriter

func init() {
	sseWriter = &SSEWriter{
		chanCache: make(map[string]chan []byte),
		mu:        &sync.Mutex{},
	}
}

// RegisterChan 注册SSE推送channel
func RegisterChan(ch chan []byte) string {
	key := fmt.Sprintf("%p-%d", ch, time.Now().UnixMilli())
	sseWriter.mu.Lock()
	defer sseWriter.mu.Unlock()
	sseWriter.chanCache[key] = ch
	return key
}

// RemoveChan 删除SSE推送channel
func RemoveChan(key string) {
	sseWriter.mu.Lock()
	defer sseWriter.mu.Unlock()
	delete(sseWriter.chanCache, key)
}

func GetSSEWriter() *SSEWriter {
	return sseWriter
}

func (s *SSEWriter) Write(p []byte) (int, error) {
	for _, ch := range s.chanCache {
		ch <- p
	}
	return 0, nil
}
