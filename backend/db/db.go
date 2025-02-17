package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load("../.env") // プロジェクトルートからの相対パスを指定
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dbURL := fmt.Sprintf("postgresql://postgres:%s@db.%s:5432/postgres?sslmode=disable",
		os.Getenv("SUPABASE_DATABASE_PASSWORD"),
		strings.Replace(os.Getenv("SUPABASE_URL"), "https://", "", 1), // https:// を削除
	)

	fmt.Println(dbURL)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
