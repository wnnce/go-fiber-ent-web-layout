package tools

import (
	"github.com/gofiber/fiber/v3"
)

func FiberRequestError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusBadRequest, message)
}

func FiberAuthError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusUnauthorized, message)
}

func FiberServerError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusInternalServerError, message)
}
