package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	err    error
	dbpool *pgxpool.Pool
)

func InitPostgresDB() {
	err = godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	os.Setenv("DATABASE_URL", "postgres://postgres:"+os.Getenv("POSTGRES_PASSWORD")+"@localhost:5432")
	dbpool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	DeletePlayersTable(ctx)
	DeleteGiftsTable(ctx)
	DeleteTagsTable(ctx)

	CreatePlayersTable(ctx)
	CreateGiftsTable(ctx)
	CreateTagsTable(ctx)
}
