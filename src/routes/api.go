package routes

import (
	"fiber_example/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/", controllers.Example)
}
