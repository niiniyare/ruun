package main

import (
    "context"
    "log"
    "os"

    "github.com/a-h/templ"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"

    "ruun-demo/sections"
)

func main() {
    app := fiber.New(fiber.Config{
        AppName: "Ruun UI - Kitchen Sink",
    })

    app.Use(logger.New())
    app.Use(recover.New())

    // Routes
    app.Get("/", render(sections.HomePage()))
    app.Get("/atoms", render(sections.AtomsPage()))
    app.Get("/molecules", render(sections.MoleculesPage()))
    app.Get("/organisms", render(sections.OrganismsPage()))
    app.Get("/templates", render(sections.TemplatesPage()))
    app.Get("/examples", render(sections.ExamplesPage()))

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Printf("🎨 Ruun Kitchen Sink: http://localhost:%s", port)
    log.Fatal(app.Listen(":" + port))
}

func render(c templ.Component) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        ctx.Set("Content-Type", "text/html")
        return c.Render(context.Background(), ctx.Response().BodyWriter())
    }
}