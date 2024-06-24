package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Success will write a default template response when returning a success response
func Success(c *fiber.Ctx, status int, data interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    status,
			"message": nil,
		},
		"data": data,
	}

	c.Set("Content-Type", "application/json")

	return c.Status(status).JSON(response)
}

// SuccessWithMessage will write a default template response when returning a success response
func SuccessWithMessage(c *fiber.Ctx, status int, data interface{}, params map[string]interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    params["code"],
			"message": params["message"],
		},
		"data": data,
	}

	c.Set("Content-Type", "application/json")

	return c.Status(status).JSON(response)
}

// SuccessOnlyMessage will write a default template response when returning a success response
func SuccessOnlyMessage(c *fiber.Ctx, status int, params map[string]interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    params["code"],
			"message": params["message"],
		},
		"data": nil,
	}

	c.Set("Content-Type", "application/json")

	return c.Status(status).JSON(response)
}

// Failed will write a default template response when returning a failed response
func Failed(c *fiber.Ctx, status int, err error) error {
	// if status/1e2 == 4 {
	// 	logger.Warn("%v", err)
	// } else {
	// 	logger.Err("%v", err)
	// }

	var errResponse map[string]interface{}

	if err != nil {
		errCode := 0
		errMsg := err.Error()
		var c *Error
		if errors.As(err, &c) {
			errCode = c.Code
			if status == 0 {
				status = c.Status
			}
			errMsg = c.Message
		}

		if errCode == 0 {
			errCode = status
		}

		errResponse = map[string]interface{}{
			"status":  status,
			"message": errMsg,
			"code":    errCode,
		}
	}

	response := errResponse

	c.Set("Content-Type", "application/json")

	return c.Status(status).JSON(response)
}
