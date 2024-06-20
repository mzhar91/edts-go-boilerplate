package middleware

import (
	"github.com/labstack/echo/v4"
)

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// origin := c.Request().Header.Get("Origin")
		// allowedOrigins := strings.Split(_config.Cfg.AllowedOrigins, ",")
		
		// for _, allowedOrigin := range allowedOrigins {
		// 	if origin == allowedOrigin {
		// 		c.Response().Header().Set("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }
		
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Response().Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Authorization, Access-Control-Requested-Token, Access-Control-Requested-For, Access-Control-Requested-Host",
		)
		
		if c.Request().Method == "OPTIONS" {
			c.Response().Write([]byte("allowed"))
		}
		
		return next(c)
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
