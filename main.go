package main

import (
	"fg/handles"

	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	handles.NewProduct(f, "this.connect-database-client")

	if err := f.Listen(":3200"); err != nil {
		panic(err)
	}
}
