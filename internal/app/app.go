package app

import (
	"log"
	"os"

	"github.com/aadejanovs/catalog/database"
	"github.com/aadejanovs/catalog/internal/app/middlewares"
	"github.com/aadejanovs/catalog/internal/app/routes"
	"github.com/aadejanovs/catalog/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/joho/godotenv"
)

func Setup() *fiber.App {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal("Mysql connection failed", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.HandleHTTPErrors,
	})

	app.Use(healthcheck.New())
	app.Use(middlewares.LoggingMiddleware())
	app.Use(middlewares.RedisMiddleware())
	app.Use(middlewares.DbMiddleware())
	app.Use(middlewares.PrometheusMiddleware())

	routes.SetupRoutes(app)
	app.Use(middlewares.CloseRedisMiddleware())

	return app
}
