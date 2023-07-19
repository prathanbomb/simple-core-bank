package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/oatsaysai/simple-core-bank/src/custom_error"
)

const APP_CTX_KEY = "appCtx"

type Response struct {
	Code       uint64      `json:"code"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func ReturnCustomError(c *fiber.Ctx, customErr error) error {
	switch customErr := customErr.(type) {
	case *custom_error.UserError:
		c.Status(http.StatusBadRequest)
		return c.JSON(&custom_error.UserError{
			Code:    customErr.Code,
			Message: customErr.Error(),
		})
	default:
		c.Status(http.StatusInternalServerError)
		return c.JSON(&custom_error.InternalError{
			Code:    custom_error.UnknownError,
			Message: customErr.Error(),
		})
	}
}
