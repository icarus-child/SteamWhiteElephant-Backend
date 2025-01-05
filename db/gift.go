package db

import (
	"context"
)

type Gift struct {
	GifterID string   `json:"gifter"`
	SteamID  int      `json:"steamId"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
}

func DeleteGiftsTable(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS gifts CASCADE;")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreateGiftsTable(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS gifts (pid UUID NOT NULL REFERENCES players ON DELETE CASCADE, steamid INTEGER NOT NULL, name TEXT, PRIMARY KEY(pid, steamid));")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetRoomGifts(c context.Context, sid string) ([]Gift, error) {
	var gifts []Gift
	rows, err := dbpool.Query(c, "SELECT * FROM gifts WHERE pid IN (SELECT p.pid FROM players p WHERE p.sid = $1);", sid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var gift Gift
		err := rows.Scan(&gift.GifterID, &gift.SteamID, &gift.Name)
		if err != nil {
			return nil, err
		}
		gift.Tags, err = GetGiftTags(c, gift.GifterID, gift.SteamID)
		if err != nil {
			return nil, err
		}
		gifts = append(gifts, gift)
	}
	return gifts, nil
}

func CreateGift(c context.Context, gift Gift) error {
	_, err := dbpool.Exec(c, "INSERT INTO gifts VALUES ($1, $2, $3);", gift.GifterID, gift.SteamID, gift.Name)
	if err != nil {
		println("hit1")
		return err
	}
	err = CreateTags(c, gift)
	if err != nil {
		println("hit2")
		return err
	}
	return nil
}
