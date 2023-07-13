package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/oatsaysai/simple-core-bank/src/app"
	"github.com/oatsaysai/simple-core-bank/src/custom_error"
	"github.com/oatsaysai/simple-core-bank/src/model"
)

func AccountRouter(fiberApp fiber.Router) {
	fiberApp.Post("/create-account", CreateAccount())
}

func CreateAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.CreateAccountParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.CreateAccount(params)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(&custom_error.InternalError{
				Code:    custom_error.UnknownError,
				Message: err.Error(),
			})
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}
