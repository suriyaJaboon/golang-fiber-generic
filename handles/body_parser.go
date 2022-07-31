package handles

import (
	"github.com/gofiber/fiber/v2"
)

// NewHandleBodyParser Generic type protect request and response.
func NewHandleBodyParser[REQUEST, RESPONSE any](handler func(REQUEST) (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Todo implement body parser OR Convert data.
		var request REQUEST
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		//Todo implement validate
		// logic validator.

		//Todo implement function handler.
		res, err := handler(request)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}

// NewHandleParamsParser Generic type
func NewHandleParamsParser[REQUEST, RESPONSE any](handler func(REQUEST) (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Todo implement body parser OR Convert data.
		var request REQUEST
		if err := c.ParamsParser(&request); err != nil {
			return err
		}

		//Todo implement validate
		// logic validator.

		//Todo implement function handler.
		res, err := handler(request)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}
