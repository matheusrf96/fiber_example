package main

import (
	"log"

	"fiber_example/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/app", controllers.Example)

	log.Fatal(app.Listen(":3000"))
}
