package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			slog.Error("Handler error", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		return nil
	}
}
