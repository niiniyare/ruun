package main

import (
	"context"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Awo ERP Component Demo",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve static assets (CSS, JS)
	app.Static("/static", "../../static/dist/")

	// Main demo page (kitchen sink)
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, DemoPage())
	})

	// Section routes for quick navigation
	app.Get("/atoms", func(c *fiber.Ctx) error {
		return render(c, AtomsSection())
	})

	app.Get("/molecules", func(c *fiber.Ctx) error {
		return render(c, MoleculesSection())
	})

	app.Get("/organisms", func(c *fiber.Ctx) error {
		return render(c, OrganismsSection())
	})

	app.Get("/templates", func(c *fiber.Ctx) error {
		return render(c, TemplatesSection())
	})

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Demo server starting on http://localhost:%s", port)
	log.Printf("📦 Kitchen Sink: http://localhost:%s/", port)
	log.Printf("⚛️  Atoms: http://localhost:%s/atoms", port)
	log.Printf("🧬 Molecules: http://localhost:%s/molecules", port)
	log.Printf("🦠 Organisms: http://localhost:%s/organisms", port)
	log.Fatal(app.Listen(":" + port))
}

// render helper for templ components
func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(context.Background(), c.Response().BodyWriter())
}
