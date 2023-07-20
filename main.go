package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! This is a test")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hi there")
	})

	app.Listen(":3000")
}
