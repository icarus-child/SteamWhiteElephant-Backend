package db

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Player struct {
	Name string
}

func DeletePlayerTable(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS players")
	if err != nil {
		return err
	}
	return nil
}

func CreatePlayerTable(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS players (id UUID PRIMARY KEY, name TEXT);")
	if err != nil {
		return err
	}
	return nil
}

func GetPlayer(c *gin.Context, id string) (*Player, error) {
	var playerId string
	var player Player
	row := dbpool.QueryRow(c, "SELECT * FROM players WHERE id = $1", id)
	err := row.Scan(&playerId, &player.Name)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func CreatePlayer(c *gin.Context, id string, player Player) error {
	_, err := dbpool.Exec(c, "INSERT INTO players VALUES ($1, $2);", id, player.Name)
	if err != nil {
		return err
	}
	return nil
}
