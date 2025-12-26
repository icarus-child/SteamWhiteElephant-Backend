package db

import (
	"context"

	"main/types"
)

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

func GetPlayerHeldGifts(c context.Context, playerid string) ([]types.Gift, error) {
	var gifts []types.Gift
	rows, err := dbpool.Query(c, "SELECT g.pid, g.steamid, g.name FROM gifts g WHERE g.pid IN (SELECT p.heldPresent FROM players p where p.pid = $1);", playerid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var gift types.Gift
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

func GetPlayerGifts(c context.Context, playerid string) ([]types.Gift, error) {
	var gifts []types.Gift
	rows, err := dbpool.Query(c, "SELECT * FROM gifts WHERE pid = $1;", playerid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var gift types.Gift
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

func GetRoomGifts(c context.Context, roomid string) ([]types.Gift, error) {
	var gifts []types.Gift
	rows, err := dbpool.Query(c, "SELECT * FROM gifts WHERE pid IN (SELECT p.pid FROM players p WHERE p.roomid = $1);", roomid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var gift types.Gift
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

func CreateGift(c context.Context, gift types.Gift) error {
	_, err := dbpool.Exec(c, "INSERT INTO gifts VALUES ($1, $2, $3);", gift.GifterID, gift.SteamID, gift.Name)
	if err != nil {
		return err
	}
	err = CreateTags(c, gift)
	if err != nil {
		return err
	}
	return nil
}
