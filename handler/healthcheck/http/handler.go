package http

import (
	"net/http"
	
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

// Healthcheck to check if the service is running
func (a *Handler) healthcheck(c echo.Context) error {
	msg := map[string]interface{}{"message": "OK"}
	
	return _api.SuccessWithMessage(c, http.StatusOK, nil, msg)
}
