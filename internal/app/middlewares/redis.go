package middlewares

import (
	"log"

	"github.com/aadejanovs/catalog/database"
	"github.com/gofiber/fiber/v2"
)

func RedisMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("redis", database.NewRedis())
		return c.Next()
	}
}

func CloseRedisMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, ok := c.Locals("redis").(*database.Redis)
		if ok {
			if err := client.Close(); err != nil {
				log.Println("Error closing Redis connection:", err)
			}
		}

		return c.Next()
	}
}
