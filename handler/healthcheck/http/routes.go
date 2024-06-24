package http

import (
	"github.com/gofiber/fiber/v2"
)

// NewHandler represent new handler
func NewHandler(e *fiber.App) {
	handler := &Handler{}
	e.Get("/healthcheck", handler.healthcheck)
}
