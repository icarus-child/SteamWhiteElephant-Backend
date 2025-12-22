package db

import (
	"context"
)

func DeleteWrappingPaper(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS wrapping CASCADE;")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreateWrappingPaper(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS wrapping (pid UUID PRIMARY KEY REFERENCES players ON DELETE CASCADE, albedo BYTEA, giftName TEXT);")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetTexture(c context.Context, pid string) ([]byte, error) {
	var texture []byte
	row := dbpool.QueryRow(c, "SELECT albedo FROM wrapping WHERE pid = $1;", pid)
	err := row.Scan(&texture)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return texture, nil
}

func GetGiftName(c context.Context, pid string) (string, error) {
	var giftName string
	row := dbpool.QueryRow(c, "SELECT giftName FROM wrapping WHERE pid = $1;", pid)
	err := row.Scan(&giftName)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return giftName, nil
}

func CreateTexture(c context.Context, pid string, data []byte) error {
	_, err := dbpool.Exec(c, "INSERT INTO wrapping (pid, albedo) VALUES ($1, $2) ON CONFLICT (pid) DO UPDATE SET albedo = EXCLUDED.albedo;", pid, data)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreateGiftName(c context.Context, pid string, giftName string) error {
	_, err := dbpool.Exec(c, "INSERT INTO wrapping (pid, giftName) VALUES ($1, $2) ON CONFLICT (pid) DO UPDATE SET giftName = EXCLUDED.giftName;", pid, giftName)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
