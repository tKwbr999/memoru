package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-github-account/memoru-backend/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/memos", handler.CreateMemo)
	api.Get("/memos", handler.GetMemos)
	api.Put("/memos/:id", handler.UpdateMemo)
	api.Delete("/memos/:id", handler.DeleteMemo)
}
