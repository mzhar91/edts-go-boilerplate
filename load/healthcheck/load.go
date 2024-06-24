package healthcheck

import (
	"github.com/gofiber/fiber/v2"

	_http "sg-edts.com/edts-go-boilerplate/handler/healthcheck/http"
)

func Load(e *fiber.App) {
	_http.NewHandler(e)
}
