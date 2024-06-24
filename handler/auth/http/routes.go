package http

import (
	"github.com/gofiber/fiber/v2"

	_credential "sg-edts.com/edts-go-boilerplate/usecase/credential"
)

// NewHandler represent new handler
func NewHandler(c *fiber.App, pu _credential.Usecase) {
	handler := &Handler{
		CredentialUseCase: pu,
	}

	g := c.Group("/auth")
	g.Post("", handler.addCredential)
	g.Post("/signin", handler.signIn)
	g.Post("/signout", handler.signOut)
	g.Get("/refresh", handler.refreshToken)
	g.Get("/availability", handler.checkTokenAvailability)
}
