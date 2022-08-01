package handles

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// NewHandleResponse Generic type protect request and response.
func NewHandleResponse[RESPONSE any](handle func() (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Todo implement function handler.
		res, err := handle()
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}

// NewHandleBodyParser Generic type protect request and response.
func NewHandleBodyParser[REQUEST, RESPONSE any](handle func(REQUEST) (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Todo implement body parser OR Convert data.
		var request REQUEST
		if err := c.BodyParser(&request); err != nil {
			log.Println(err)
			return ErrBodyParser()
		}

		//Todo implement validate
		// logic validator.
		if err := Validation(request); len(err) > 0 {
			return ErrValidator(err)
		}

		//Todo implement function handler.
		res, err := handle(request)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}

// NewHandleParamsParser Generic type
func NewHandleParamsParser[REQUEST, RESPONSE any](handle func(REQUEST) (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Todo implement params parser OR Convert data.
		var request REQUEST
		if err := c.ParamsParser(&request); err != nil {
			log.Println(err)
			return ErrBodyParser()
		}

		//Todo implement validate
		// logic validator.
		if err := Validation(request); len(err) > 0 {
			return ErrValidator(err)
		}

		//Todo implement function handler.
		res, err := handle(request)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}

// NewHandleParamsWithBodyParser Generic type
func NewHandleParamsWithBodyParser[PARAMS, REQUEST, RESPONSE any](handle func(PARAMS, REQUEST) (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Todo implement params parser OR Convert data.
		var id PARAMS
		if err := c.ParamsParser(&id); err != nil {
			return err
		}

		//Todo implement body parser OR Convert data.
		var request REQUEST
		if err := c.BodyParser(&request); err != nil {
			log.Println(err)
			return ErrBodyParser()
		}

		//Todo implement validate
		// logic validator.
		if err := Validation(request); len(err) > 0 {
			return ErrValidator(err)
		}

		//Todo implement function handler.
		res, err := handle(id, request)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}
