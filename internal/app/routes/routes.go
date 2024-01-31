package routes

import (
	"github.com/aadejanovs/catalog/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/blueprints/:id", handlers.GetBlueprint)
	app.Get("/blueprints", handlers.ListBlueprints)
	app.Post("/blueprints", handlers.CreateBlueprint)
	app.Get("/metrics", handlers.GetMetrics(promhttp.Handler()))

	// Endpoint used for testing purposes.
	app.Get("/cursor-blueprints", handlers.CursorListBlueprints)
}
