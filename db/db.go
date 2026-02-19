package db

import (
	"context"
	"fmt"
	"main/types"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	err    error
	dbpool *pgxpool.Pool
)

func InitPostgresDB() {
	println("postgres://postgres:" + os.Getenv("POSTGRES_PASSWORD") + "@postgres:5432")
	os.Setenv("DATABASE_URL", "postgres://postgres:"+os.Getenv("POSTGRES_PASSWORD")+"@postgres:5432")
	dbpool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	DeletePlayersTable(ctx)
	DeleteGiftsTable(ctx)
	DeleteTagsTable(ctx)
	DeleteRoomTable(ctx)
	DeleteWrappingPaper(ctx)

	CreatePlayersTable(ctx)
	CreateGiftsTable(ctx)
	CreateTagsTable(ctx)
	CreateRoomTable(ctx)
	CreateWrappingPaper(ctx)
}

type Room struct {
	Id       string              `json:"id"`
	Players  []Player            `json:"players"`
	Presents []types.PresentJson `json:"presents"`
}
