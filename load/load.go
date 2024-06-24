package load

import (
	"time"

	"github.com/gofiber/fiber/v2"

	_config "sg-edts.com/edts-go-boilerplate/config"
	_loadAuth "sg-edts.com/edts-go-boilerplate/load/credential"
	_loadHealth "sg-edts.com/edts-go-boilerplate/load/healthcheck"
)

func Load(e *fiber.App, connection *_config.Connection, timeoutContext time.Duration) {
	_loadAuth.Load(e, connection, timeoutContext)
	_loadHealth.Load(e)
}
