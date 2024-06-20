package api

import (
	"errors"
	
	"github.com/labstack/echo/v4"
)

// Success will write a default template response when returning a success response
func Success(c echo.Context, status int, data interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    status,
			"message": nil,
		},
		"data": data,
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}

// SuccessWithMessage will write a default template response when returning a success response
func SuccessWithMessage(c echo.Context, status int, data interface{}, params map[string]interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    params["code"],
			"message": params["message"],
		},
		"data": data,
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}

// SuccessOnlyMessage will write a default template response when returning a success response
func SuccessOnlyMessage(c echo.Context, status int, params map[string]interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    params["code"],
			"message": params["message"],
		},
		"data": nil,
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}

// Failed will write a default template response when returning a failed response
func Failed(c echo.Context, status int, err error) error {
	// if status/1e2 == 4 {
	// 	logger.Warn("%v", err)
	// } else {
	// 	logger.Err("%v", err)
	// }
	
	var errResponse map[string]interface{}
	
	if err != nil {
		errCode := 0
		errMsg := err.Error()
		var f *Error
		if errors.As(err, &f) {
			errCode = f.Code
			if status == 0 {
				status = f.Status
			}
			errMsg = f.Message
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
	
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}
