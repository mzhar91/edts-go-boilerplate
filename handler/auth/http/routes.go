package http

import (
	"github.com/labstack/echo/v4"
	
	_credential "sg-edts.com/edts-go-boilerplate/usecase/credential"
)

// NewHandler represent new handler
func NewHandler(e *echo.Echo, pu _credential.Usecase) {
	handler := &Handler{
		CredentialUseCase: pu,
	}
	
	g := e.Group("/auth")
	g.POST("", handler.addCredential)
	g.POST("/signin", handler.signIn)
	g.POST("/signout", handler.signOut)
	g.GET("/refresh", handler.refreshToken)
	g.GET("/availability", handler.checkTokenAvailability)
}
