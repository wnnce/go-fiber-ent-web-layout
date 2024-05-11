package hand

import (
	"sync"
	"sync/atomic"
)

type ErrorHandlerResult struct {
	Code    int    // http请求返回响应码
	Message string // 返回的错误消息
}

// ErrorHandler 错误处理函数
type ErrorHandler func(err error) *ErrorHandlerResult

// 错误处理链
type errorHandlerChain struct {
	handlers  []ErrorHandler // 错误处理函数切片
	lastIndex int32          // 上次处理成功的错误处理函数索引
	mutex     sync.Mutex     // 锁
}

// RegisterErrorHandler 注册错误处理函数
func (e *errorHandlerChain) RegisterErrorHandler(handler ErrorHandler) *errorHandlerChain {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.handlers = append(e.handlers, handler)
	return e
}

// DoHandler 通过处理链 每个注册的错误函数 直到传入的错误被处理或者全部全部执行完成
func (e *errorHandlerChain) DoHandler(err error) *ErrorHandlerResult {
	lastIndex := atomic.LoadInt32(&e.lastIndex)
	if lastIndex > 0 {
		if result := e.handlers[lastIndex](err); result != nil {
			return result
		}
	}
	for i, handler := range e.handlers {
		if result := handler(err); result != nil {
			if i != 0 {
				atomic.StoreInt32(&e.lastIndex, int32(i))
			}
			return result
		}
	}
	return nil
}
