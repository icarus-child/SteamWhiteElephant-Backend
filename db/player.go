package db

import (
	"context"
)

type Player struct {
	PlayerID string `json:"playerId"`
	RoomID   string `json:"roomId"`
	Name     string `json:"name"`
}

func DeletePlayersTable(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS players CASCADE;")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreatePlayersTable(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS players (pid UUID PRIMARY KEY, roomid TEXT NOT NULL, name TEXT);")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetPlayer(c context.Context, pid string) (*Player, error) {
	var player Player = Player{PlayerID: pid}
	row := dbpool.QueryRow(c, "SELECT roomid, name FROM players WHERE pid = $1;", pid)
	err := row.Scan(&player.RoomID, &player.Name)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func GetRoomPlayers(c context.Context, roomid string) ([]Player, error) {
	var players []Player
	rows, err := dbpool.Query(c, "SELECT pid, name FROM players WHERE roomid = $1;", roomid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var player Player = Player{RoomID: roomid}
		err := rows.Scan(&player.PlayerID, &player.Name)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func CheckRoomExists(c context.Context, roomid string) (bool, error) {
	var ret bool
	row := dbpool.QueryRow(c, "SELECT EXISTS (SELECT * FROM players WHERE roomid = $1);", roomid)
	err := row.Scan(&ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func CreatePlayer(c context.Context, pid string, roomid string, player Player) error {
	_, err := dbpool.Exec(c, "INSERT INTO players VALUES ($1, $2, $3);", pid, roomid, player.Name)
	if err != nil {
		return err
	}
	return nil
}
