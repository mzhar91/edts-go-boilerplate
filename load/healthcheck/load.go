package healthcheck

import (
	"github.com/labstack/echo/v4"
	
	_http "sg-edts.com/edts-go-boilerplate/handler/healthcheck/http"
)

func Load(e *echo.Echo) {
	_http.NewHandler(e)
}
