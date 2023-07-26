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
	fiberApp.Post("/get-account", GetAccount())
	fiberApp.Post("/pre-generate-account-no", PreGenerateAccountNo())
	fiberApp.Post("/get-transaction", GetTransaction())
}

func PreGenerateAccountNo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.PreGenerateAccountNoParams

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
		_, err = appCtx.PreGenerateAccountNumbers(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "Successfully pre-generated account numbers",
			Data:    nil,
		})
	}
}

func GetAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.GetAccountParams

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
		result, err := appCtx.GetAccount(&params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
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
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}

func GetTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.GetTransactionParams

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
		result, err := appCtx.GetTransactionByAccountNo(&params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}
