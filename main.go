package main

import (
	"logistica/app/routes"
	"logistica/app/utilities"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	utilities.ClearScreen()                        // Membersihkan Console
	engine := mustache.New("./views", ".mustache") // Inisialisasi Mustache sebagai Template Engine
	app := fiber.New(fiber.Config{
		Views: engine,
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return c.Redirect("/404")
		// },
	})
	app.Static("/", "./public") // Konfigurasi mdlwr untuk file statis seperti css, js
	// Routes
	routes.SetupRoutes(app)
	// Menjalankan Server di Port 9999
	log.Fatal(app.Listen(":9999"))
}
