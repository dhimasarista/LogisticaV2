package routes

import "github.com/gofiber/fiber/v2"

func InventarisRoutes(app *fiber.App) {
	// Render Halaman Inventaris
	app.Get("/inventaris", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Render("inventaris", fiber.Map{
			"path":     path,
			"username": "dhimasarista",
		}, "layouts/main")
	})
}
