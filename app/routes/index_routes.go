package routes

import "github.com/gofiber/fiber/v2"

func IndexRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/dashboard")
	})
}
