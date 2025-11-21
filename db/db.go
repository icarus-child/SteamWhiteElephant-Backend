package db

import (
	"context"
	"fmt"
	"main/types"
	"main/utility"
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

type Room struct {
	Id       string              `json:"id"`
	Players  []Player            `json:"players"`
	Presents []types.PresentJson `json:"presents"`
}

func GetAll(ctx context.Context) (ret []Room, err error) {
	rooms, err := dbpool.Query(ctx, "SELECT DISTINCT roomid FROM players;")
	if err != nil {
		return nil, err
	}
	for rooms.Next() {
		var room Room
		err := rooms.Scan(&room.Id)
		if err != nil {
			return nil, err
		}
		room.Players, err = GetRoomPlayers(ctx, room.Id)
		if err != nil {
			return nil, err
		}
		tempPresents, err := GetRoomGifts(ctx, room.Id)
		if err != nil {
			return nil, err
		}
		room.Presents = utility.CollectPresentsByGifter(tempPresents)
		ret = append(ret, room)
	}
	return
}
