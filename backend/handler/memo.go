package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/your-github-account/memoru-backend/db"
	"github.com/your-github-account/memoru-backend/model"
)

func CreateMemo(c *fiber.Ctx) error {
	dbConn, err := db.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer dbConn.Close()

	var memo model.Memo
	if err := c.BodyParser(&memo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := dbConn.Exec("INSERT INTO memos (content) VALUES ($1)", memo.Content)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Error inserting memo: %v", err)})
	}

	// InsertされたレコードのIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Error getting last insert ID: %v", err)})
	}

	memo.ID = fmt.Sprintf("%d", id) // 取得したIDをmemoにセット

	return c.Status(http.StatusCreated).JSON(memo)
}

func GetMemos(c *fiber.Ctx) error {
	dbConn, err := db.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT id, content, created_at FROM memos")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var memos []model.Memo
	for rows.Next() {
		var memo model.Memo
		if err := rows.Scan(&memo.ID, &memo.Content, &memo.CreatedAt); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		memos = append(memos, memo)
	}

	return c.JSON(memos)
}

func UpdateMemo(c *fiber.Ctx) error {
	dbConn, err := db.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer dbConn.Close()

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "memo ID is required"})
	}

	var memo model.Memo
	if err := c.BodyParser(&memo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = dbConn.Exec("UPDATE memos SET content = $1 WHERE id = $2", memo.Content, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(memo)
}

func DeleteMemo(c *fiber.Ctx) error {
	dbConn, err := db.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer dbConn.Close()

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "memo ID is required"})
	}

	_, err = dbConn.Exec("DELETE FROM memos WHERE id = $1", id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
