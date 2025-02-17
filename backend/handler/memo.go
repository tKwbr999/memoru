package handler

import (
	"database/sql"
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/your-github-account/memoru-backend/model"
)

func CreateMemo(c *fiber.Ctx) error {
	dbConn, ok := c.Locals("db").(*sql.DB)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get database connection from context"})
	}

	var memo model.Memo
	if err := c.BodyParser(&memo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errors.Wrap(err, "failed to parse request body")})
	}

	var id string
	err := dbConn.QueryRow("INSERT INTO memos (content) VALUES ($1) RETURNING id", memo.Content).Scan(&id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.Wrapf(err, "failed to insert memo: %s:%d", file, line)})
	}

	memo.ID = id // 取得したIDをmemoにセット

	return c.Status(http.StatusCreated).JSON(memo)
}

func GetMemos(c *fiber.Ctx) error {
	dbConn, ok := c.Locals("db").(*sql.DB)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get database connection from context"})
	}

	rows, err := dbConn.Query("SELECT id, content, created_at FROM memos")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.Wrapf(err, "failed to query memos: %s:%d", file, line)})
	}

	var memos []model.Memo
	for rows.Next() {
		var memo model.Memo
		if err := rows.Scan(&memo.ID, &memo.Content, &memo.CreatedAt); err != nil {
			_, file, line, _ := runtime.Caller(0)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.Wrapf(err, "failed to scan memo: %s:%d", file, line)})
		}
		memos = append(memos, memo)
	}

	return c.JSON(memos)
}

func UpdateMemo(c *fiber.Ctx) error {
	dbConn, ok := c.Locals("db").(*sql.DB)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get database connection from context"})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "memo ID is required"})
	}

	var memo model.Memo
	if err := c.BodyParser(&memo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errors.Wrap(err, "failed to parse request body")})
	}

	_, err := dbConn.Exec("UPDATE memos SET content = $1 WHERE id = $2", memo.Content, id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.Wrapf(err, "failed to update memo: %s:%d", file, line)})
	}

	return c.JSON(memo)
}

func DeleteMemo(c *fiber.Ctx) error {
	dbConn, ok := c.Locals("db").(*sql.DB)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get database connection from context"})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "memo ID is required"})
	}

	_, err := dbConn.Exec("DELETE FROM memos WHERE id = $1", id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.Wrapf(err, "failed to delete memo: %s:%d", file, line)})
	}

	return c.SendStatus(http.StatusNoContent)
}
