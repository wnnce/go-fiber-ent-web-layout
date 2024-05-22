package res

import (
	"github.com/gofiber/fiber/v3"
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
		Code:      fiber.StatusOK,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}

func OkByMessage(message string) *Result[any] {
	return &Result[any]{
		Code:      fiber.StatusOK,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
	}
}

func OkByData[T any](data T) *Result[T] {
	return &Result[T]{
		Code:      fiber.StatusOK,
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
