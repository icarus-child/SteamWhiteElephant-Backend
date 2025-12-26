package db

import (
	"context"
)

type Player struct {
	PlayerID  string  `json:"playerId"`
	RoomID    string  `json:"roomId"`
	Name      string  `json:"name"`
	PresentID *string `json:"presentId"`
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
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS players (pid UUID PRIMARY KEY, roomid TEXT NOT NULL, name TEXT, heldPresent UUID REFERENCES players(pid));")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetPlayer(c context.Context, pid string) (*Player, error) {
	var player Player = Player{PlayerID: pid}
	row := dbpool.QueryRow(c, "SELECT roomid, name, heldPresent FROM players WHERE pid = $1;", pid)
	err := row.Scan(&player.RoomID, &player.Name, &player.PresentID)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func GetRoomPlayers(c context.Context, roomid string) ([]Player, error) {
	var players []Player
	rows, err := dbpool.Query(c, "SELECT pid, name, heldPresent FROM players WHERE roomid = $1;", roomid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var player Player = Player{RoomID: roomid}
		err := rows.Scan(&player.PlayerID, &player.Name, &player.PresentID)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func GetOrderedRoomPlayers(c context.Context, roomID string) ([]Player, error) {
	var players []Player

	rows, err := dbpool.Query(
		c,
		`SELECT p.pid, p.name, p.heldPresent
		 FROM room r
		 JOIN unnest(r.playerOrder) WITH ORDINALITY AS o(pid, ord) ON TRUE
		 JOIN players p ON p.pid = o.pid
		 WHERE r.rid = $1
		 ORDER BY o.ord`,
		roomID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var player Player
		player.RoomID = roomID

		err := rows.Scan(&player.PlayerID, &player.Name, &player.PresentID)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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

func CreatePlayer(c context.Context, pid string, roomID string, player Player) error {
	// 1️⃣ Ensure room exists
	_, err := dbpool.Exec(
		c,
		"INSERT INTO room (rid, started) VALUES ($1, false) ON CONFLICT DO NOTHING",
		roomID,
	)
	if err != nil {
		return err
	}

	// 2️⃣ Insert player
	_, err = dbpool.Exec(
		c,
		"INSERT INTO players (pid, roomid, name) VALUES ($1, $2, $3);",
		pid, roomID, player.Name,
	)
	if err != nil {
		return err
	}

	// 3️⃣ Append player ID to room's playerOrder
	_, err = dbpool.Exec(
		c,
		"UPDATE room SET playerOrder = array_append(playerOrder, $2::UUID) WHERE rid = $1",
		roomID, pid,
	)
	if err != nil {
		return err
	}

	return nil
}

func TakeOrStealPresent(c context.Context, playerID string, presentID string) error {
	// 1️⃣ Remove the present from whoever currently holds it
	_, err := dbpool.Exec(c, `
		UPDATE players
		SET heldPresent = NULL
		WHERE heldPresent = $1
	`, presentID)
	if err != nil {
		return err
	}

	// 2️⃣ Assign the present to the new player
	_, err = dbpool.Exec(c, `
		UPDATE players
		SET heldPresent = $1
		WHERE pid = $2
	`, presentID, playerID)
	return err
}
