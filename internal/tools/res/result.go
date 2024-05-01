package res

import (
	"net/http"
	"time"
)

type Result[T any] struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      T      `json:"data,omitempty"`
}

func Ok[T any](message string, data T) *Result[T] {
	return &Result[T]{
		Code:      http.StatusOK,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}

func OkByMessage(message string) *Result[any] {
	return &Result[any]{
		Code:      http.StatusOK,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
	}
}

func OkByData[T any](data T) *Result[T] {
	return &Result[T]{
		Code:      http.StatusOK,
		Message:   "ok",
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}

func Fail(code int, message string) *Result[any] {
	return &Result[any]{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
	}
}
