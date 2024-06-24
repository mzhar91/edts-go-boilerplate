package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
)

type Handler struct {
}

// @Tags Health Check
// @Summary Health Check
// @Description Health check endpoint
// @Produce  json
// @Router /healthcheck [get]
func (a *Handler) healthcheck(c *fiber.Ctx) error {
	msg := map[string]interface{}{"message": "OK"}

	return _api.SuccessWithMessage(c, http.StatusOK, nil, msg)
}
