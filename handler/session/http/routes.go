package http

import (
	"github.com/gofiber/fiber/v2"

	_session "sg-edts.com/edts-go-boilerplate/usecase/session"
)

// NewHandler represent new handler
func NewHandler(e *fiber.Ctx, pu _session.Usecase) {
	handler := &Handler{
		SessionUseCase: pu,
	}

	g := e.Group("/session")
	g.DELETE("/", handler.dropOwnSession)
	g.DELETE("/:username", handler.dropSession)
	g.GET("/", handler.getOwnSession)
	g.GET("/all", handler.getSession)
}
