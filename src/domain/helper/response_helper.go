package helper

import (
	"net/http"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/model"

	"github.com/gofiber/fiber/v2"
)

var MessageOK string

func ResponseExit(httpCode int, message string) {
	exception.PanicIfNeeded(exception.ValidationError{Status: httpCode, Message: message})
}

func ResponseError(httpCode int, message string, errorCode int) {
	exception.PanicIfNeeded(exception.ValidationError{Status: httpCode, Message: message, ErrorCode: errorCode})
}

func Response400(message string, errorCode int) {
	exception.PanicIfNeeded(exception.ValidationError{Status: 400, Message: message, ErrorCode: errorCode})
}

func ResponseLog400(message string, errorCode int) {
	// logger.Error().Msg(message)
	exception.PanicIfNeeded(exception.ValidationError{Status: 400, Message: message, ErrorCode: errorCode})
}

func Response401(message string, errorCode int) {
	exception.PanicIfNeeded(exception.ValidationError{Status: 401, Message: message, ErrorCode: errorCode})
}

func Response404(message string, errorCode int) {
	exception.PanicIfNeeded(exception.ValidationError{Status: 404, Message: message, ErrorCode: errorCode})
}

func ResponseOK(c *fiber.Ctx, response interface{}) error {

	return c.Status(http.StatusOK).JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Message: MessageOK,
		Data:    response,
	})
}

func ResponseNotFound(c *fiber.Ctx, response interface{}) error {

	return c.Status(http.StatusNotFound).JSON(model.WebResponse{
		Code:    fiber.StatusNotFound,
		Message: MessageOK,
		Data:    response,
	})
}
