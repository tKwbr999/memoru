package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load("../.env") // プロジェクトルートからの相対パスを指定
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// //transaction pooler
	// dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?pgbouncer=true&connection_limit=1",
	// 	os.Getenv("SUPABASE_DATABASE_USER"),
	// 	os.Getenv("SUPABASE_DATABASE_PASSWORD"),
	// 	os.Getenv("SUPABASE_DATABASE_HOST"),
	// 	os.Getenv("SUPABASE_DATABASE_TRANSACTION_POOLER_PORT"),
	// 	os.Getenv("SUPABASE_DATABASE_DBNAME"))

	// // session pooler
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?pgbouncer=true&connection_limit=1",
		os.Getenv("SUPABASE_DATABASE_USER"),
		os.Getenv("SUPABASE_DATABASE_PASSWORD"),
		os.Getenv("SUPABASE_DATABASE_HOST"),
		os.Getenv("SUPABASE_DATABASE_SESSION_POOLER_PORT"),
		os.Getenv("SUPABASE_DATABASE_DBNAME"))

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
