package tools

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func FiberRequestError(message string) *fiber.Error {
	return fiber.NewError(http.StatusBadRequest, message)
}

func FiberAuthError(message string) *fiber.Error {
	return fiber.NewError(http.StatusUnauthorized, message)
}

func FiberServerError(message string) *fiber.Error {
	return fiber.NewError(http.StatusInternalServerError, message)
}
