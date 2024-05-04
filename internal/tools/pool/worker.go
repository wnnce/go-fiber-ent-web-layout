package pool

import (
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
)

var workerPool sync.Pool

type worker struct {
	pool *simplePool
}

func init() {
	workerPool.New = func() any {
		return &worker{}
	}
}

func (w *worker) close() {
	w.pool.subWorkCount()
}
func (w *worker) Recycle() {
	w.pool = nil
	workerPool.Put(w)
}

// worker协程运行
// 通过for循环一直遍历任务链表 直至全部执行完将当前worker放回对象池
// 当存在多个worker执行时 对任务链表的执行需要加锁 确保任务不会重复执行
func (w *worker) run() {
	go func() {
		for {
			var t *task
			w.pool.taskLock.Lock()
			if w.pool.taskHead == nil {
				w.close()
				w.pool.taskLock.Unlock()
				w.Recycle()
				return
			} else {
				t = w.pool.taskHead
				w.pool.taskHead = w.pool.taskHead.next
				atomic.AddInt32(&w.pool.taskCount, -1)
			}
			w.pool.taskLock.Unlock()
			// 嵌套一个 func 用于捕获panic
			// 如果在当前 func 内执行或中断for循环 可能造成剩余的任务无法执行
			func() {
				defer func() {
					if r := recover(); r != nil {
						if t.handler != nil {
							t.handler(t.ctx, r)
						} else {
							msg := fmt.Sprintf("pool run panic name:%s, error:%v", w.pool.Name(), r)
							slog.ErrorContext(t.ctx, msg)
						}
					}
				}()
				t.fn()
			}()
			t.Recycle()
		}
	}()
}
