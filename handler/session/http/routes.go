package http

import (
	"github.com/labstack/echo/v4"
	
	_session "sg-edts.com/edts-go-boilerplate/usecase/session"
)

// NewHandler represent new handler
func NewHandler(e *echo.Echo, pu _session.Usecase) {
	handler := &Handler{
		SessionUseCase: pu,
	}
	
	g := e.Group("/session")
	g.DELETE("/", handler.dropOwnSession)
	g.DELETE("/:username", handler.dropSession)
	g.GET("/", handler.getOwnSession)
	g.GET("/all", handler.getSession)
}
