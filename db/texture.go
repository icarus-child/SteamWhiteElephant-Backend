package db

import (
	"context"
)

func DeleteTextureTable(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS texture CASCADE;")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreateTextureTable(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS texture (pid UUID PRIMARY KEY REFERENCES players ON DELETE CASCADE, albedo BYTEA NOT NULL);")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetTexture(c context.Context, pid string) ([]byte, error) {
	var texture []byte
	row := dbpool.QueryRow(c, "SELECT albedo FROM texture WHERE pid = $1;", pid)
	err := row.Scan(&texture)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return texture, nil
}

func CreateTexture(c context.Context, pid string, data []byte) error {
	_, err := dbpool.Exec(c, "INSERT INTO texture VALUES ($1, $2);", pid, data)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
