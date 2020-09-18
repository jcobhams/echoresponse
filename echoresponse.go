package echoresponse

import "github.com/labstack/echo/v4"

type (
	r struct {
		Message interface{} `json:"msg"`
	}
	rWithPayload struct {
		r
		Payload interface{} `json:"payload"`
	}
	Payload map[string]interface{}
)

func Format(c echo.Context, message interface{}, payload interface{}, statusCode int) error {
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

func FormatResponse(c echo.Context, message interface{}, payload interface{}, statusCode int) error {
	return Format(c, message, payload, statusCode)
}
