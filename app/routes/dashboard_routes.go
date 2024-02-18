package routes

import "github.com/gofiber/fiber/v2"

func DashboardRoutes(app *fiber.App) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path = c.Path()
		return c.Render("dashboard", fiber.Map{
			"path":     path,
			"title":    "Dashboard",
			"username": "dhimasarista",
		}, "layouts/main")
	})
}
