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
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS room (rid TEXT PRIMARY KEY, started BOOLEAN);")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetRoom(c context.Context, roomid string) (*NarrowRoom, error) {
	var room NarrowRoom
	row := dbpool.QueryRow(c, "SELECT * FROM room WHERE rid = $1;", roomid)
	err := row.Scan(&room.RoomID, &room.Started)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func CreateRoom(c context.Context, room NarrowRoom) error {
	_, err := dbpool.Exec(c, "INSERT INTO room VALUES ($1, $2);", room.RoomID, room.Started)
	if err != nil {
		return err
	}
	return nil
}
