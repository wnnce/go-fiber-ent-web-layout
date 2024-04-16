package res

import (
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
