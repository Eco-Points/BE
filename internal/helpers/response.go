package helpers

import "github.com/labstack/echo/v4"

func ResponseFormat(code int, status string, message string, data any) map[string]any {
	var result = make(map[string]any)

	result["code"] = code
	result["message"] = message
	result["status"] = status
	if data != nil {
		result["data"] = data
	}

	return result
}

func EasyHelper(c echo.Context, code int, status string, message string, data any) error {
	return c.JSON(code, ResponseFormat(code, status, message, data))

}
