package http

import (
	"github.com/labstack/echo/v4"
)

// NewHandler represent new handler
func NewHandler(e *echo.Echo) {
	handler := &Handler{}
	e.GET("/healthcheck", handler.healthcheck)
}
