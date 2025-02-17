package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tKwbr999/memoru-backend/middleware"
	"github.com/tKwbr999/memoru-backend/router"
)

func main() {
	// slogのデフォルトロガーを設定
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug, // すべてのレベルのログを出力
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, opts)))

	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.ErrorHandler()) // エラーハンドリングミドルウェアを追加

	router.SetupRoutes(app)

	app.Listen(":8080")
}
