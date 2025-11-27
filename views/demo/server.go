package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:               "Awo ERP Component Demo",
		DisableStartupMessage: false,
		ErrorHandler:          customErrorHandler,
	})

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New())

	// Serve static assets (CSS, JS)
	app.Static("/static", "../../static/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        false,
		CacheDuration: 86400, // 24 hours
	})

	// Routes
	setupRoutes(app)

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Demo server starting on http://localhost:%s", port)
	log.Printf("📦 Component Gallery: http://localhost:%s/", port)
	log.Printf("💚 Health Check: http://localhost:%s/health", port)
	log.Fatal(app.Listen(":" + port))
}

// setupRoutes configures all application routes
func setupRoutes(app *fiber.App) {
	// Main demo page (kitchen sink)
	app.Get("/", handleDemoPage)

	// Individual component demos (optional - for testing specific components)
	app.Get("/atoms", handleAtomsDemo)
	app.Get("/molecules", handleMoleculesDemo)
	app.Get("/organisms", handleOrganismsDemo)

	// Health check endpoint
	app.Get("/health", handleHealthCheck)

	// API routes group (for future HTMX endpoints)
	api := app.Group("/api")
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})
}

// handleDemoPage renders the complete component gallery
func handleDemoPage(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return DemoPage().Render(context.Background(), c.Response().BodyWriter())
}

// handleAtomsDemo renders only the atoms section
func handleAtomsDemo(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return AtomsSection().Render(context.Background(), c.Response().BodyWriter())
}

// handleMoleculesDemo renders only the molecules section
func handleMoleculesDemo(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return MoleculesSection().Render(context.Background(), c.Response().BodyWriter())
}

// handleOrganismsDemo renders only the organisms section
func handleOrganismsDemo(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return OrganismsSection().Render(context.Background(), c.Response().BodyWriter())
}

// handleHealthCheck returns service health status
func handleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": "awo-erp-demo",
		"version": "1.0.0",
	})
}

// customErrorHandler handles application errors
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	log.Printf("Error: %v", err)

	// Return JSON for API routes
	if len(c.Path()) >= 4 && c.Path()[:4] == "/api" {
		return c.Status(code).JSON(fiber.Map{
			"error":  message,
			"code":   code,
			"path":   c.Path(),
			"method": c.Method(),
		})
	}

	// Return HTML for regular routes
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Status(code).SendString(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Error ` + string(rune(code)) + `</title>
			<style>
				body { font-family: system-ui; text-align: center; padding: 2rem; }
				h1 { color: #e11d48; }
			</style>
		</head>
		<body>
			<h1>` + message + `</h1>
			<p>Error Code: ` + string(rune(code)) + `</p>
			<a href="/">← Back to Home</a>
		</body>
		</html>
	`)
}
