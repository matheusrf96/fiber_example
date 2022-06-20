package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"id":     5,
			"active": true,
			"msg":    "Show!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
