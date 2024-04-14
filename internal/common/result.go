package common

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
	"time"
)

type Result struct {
	Code      int         `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func Ok(message string, data interface{}) *Result {
	return &Result{
		Code:      http.StatusOK,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}

func OkByMessage(message string) *Result {
	return &Result{
		Code:      http.StatusOK,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
	}
}

func OkByData(data interface{}) *Result {
	return &Result{
		Code:      http.StatusOK,
		Message:   "ok",
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}

func Fail(code int, message string) *Result {
	return &Result{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
	}
}

func FiberRequestError(message string) *fiber.Error {
	return fiber.NewError(http.StatusBadRequest, message)
}

func FiberAuthError(message string) *fiber.Error {
	return fiber.NewError(http.StatusUnauthorized, message)
}

func FiberServerError(message string) *fiber.Error {
	return fiber.NewError(http.StatusInternalServerError, message)
}
