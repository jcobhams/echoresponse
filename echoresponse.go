package echoresponse

import "github.com/labstack/echo/v4"

type (
	r struct {
		Message any `json:"msg"`
	}
	rWithPayload struct {
		r
		Payload any `json:"payload"`
	}
	Payload map[string]any
)

func Format(c echo.Context, message any, payload any, statusCode int) error {
	var responseMsg string
	switch msg := message.(type) {
	case error:
		responseMsg = msg.Error()
	case string:
		responseMsg = msg
	}

	response := &r{Message: responseMsg}
	if payload != nil {
		response := &rWithPayload{r: *response, Payload: payload}
		return c.JSON(statusCode, response)
	}
	return c.JSON(statusCode, response)
}

// DEPRECATED: use Format instead
func FormatResponse(c echo.Context, message any, payload any, statusCode int) error {
	return Format(c, message, payload, statusCode)
}
