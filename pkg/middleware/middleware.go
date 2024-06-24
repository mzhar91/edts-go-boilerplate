package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *GoMiddleware) CORS(c *fiber.Ctx) error {
	// origin := c.Request().Header.Get("Origin")
	// allowedOrigins := strings.Split(_config.Cfg.AllowedOrigins, ",")

	// for _, allowedOrigin := range allowedOrigins {
	// 	if origin == allowedOrigin {
	// 		c.Response().Header().Set("Access-Control-Allow-Origin", origin)
	// 		break
	// 	}
	// }

	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	c.Set(
		"Access-Control-Allow-Headers",
		"Content-Type, Access-Control-Allow-Headers, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Authorization, Access-Control-Requested-Token, Access-Control-Requested-For, Access-Control-Requested-Host",
	)

	if c.Method() == "OPTIONS" {
		_, err := c.Write([]byte("allowed"))
		if err != nil {
			return err
		}
	}

	return c.Next()
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
