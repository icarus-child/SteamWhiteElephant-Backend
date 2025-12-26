package db

import (
	"context"
)

type NarrowRoom struct {
	RoomID  string `json:"roomId"`
	Started bool   `json:"started"`
}

func DeleteRoomTable(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS room CASCADE;")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreateRoomTable(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS room (rid TEXT PRIMARY KEY, started BOOLEAN, turnIndex SMALLINT NOT NULL DEFAULT 0, playerOrder UUID[] NOT NULL DEFAULT '{}'::UUID[]);")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetRoom(c context.Context, roomid string) (*NarrowRoom, error) {
	var room NarrowRoom
	row := dbpool.QueryRow(c, "SELECT rid, started FROM room WHERE rid = $1;", roomid)
	err := row.Scan(&room.RoomID, &room.Started)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func StartRoom(c context.Context, roomid string) error {
	_, err := dbpool.Exec(
		c,
		`UPDATE room
		 SET started = TRUE
		 WHERE rid = $1`,
		roomid,
	)
	return err
}

func GetRoomTurnIndex(c context.Context, roomid string) (int16, error) {
	var index int16
	row := dbpool.QueryRow(c, "SELECT turnIndex FROM room WHERE rid = $1;", roomid)
	err := row.Scan(&index)
	if err != nil {
		return 0, err
	}
	return index, nil
}

func SetRoomTurnIndex(c context.Context, roomid string, index int16) error {
	_, err := dbpool.Exec(c, "UPDATE room SET turnIndex = $1 WHERE rid = $2;", index, roomid)
	if err != nil {
		return err
	}
	return nil
}

func RandomizePlayerOrder(c context.Context, roomid string) error {
	_, err := dbpool.Exec(
		c,
		`UPDATE room
		SET playerOrder = (
			SELECT array_agg(pid ORDER BY random())
			FROM unnest(playerOrder) AS t(pid)
		)
		WHERE rid = $1;`,
		roomid,
	)
	return err
}
