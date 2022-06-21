package main

import (
	"log"

	"fiber_example/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Register(app)

	log.Fatal(app.Listen(":3000"))
}
