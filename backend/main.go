package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/your-github-account/memoru-backend/db"
	"github.com/your-github-account/memoru-backend/middleware"
	"github.com/your-github-account/memoru-backend/router"
)

func main() {
	// slogのデフォルトロガーを設定
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	// DBコネクションを取得
	dbConn, err := db.Connect()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1) // 接続に失敗したらプログラムを終了
	}
	defer dbConn.Close()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.ErrorHandler()) // エラーハンドリングミドルウェアを追加

	// DBコネクションをContextに保存するミドルウェア
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConn)
		return c.Next()
	})

	router.SetupRoutes(app)

	app.Listen(":8080")
}
