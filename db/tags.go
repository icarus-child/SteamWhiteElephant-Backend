package db

import "context"

func DeleteTagsTable(c context.Context) error {
	_, err := dbpool.Exec(c, "DROP TABLE IF EXISTS tags CASCADE;")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CreateTagsTable(c context.Context) error {
	_, err := dbpool.Exec(c, "CREATE TABLE IF NOT EXISTS tags (pid UUID NOT NULL, steamid INTEGER NOT NULL, tag TEXT NOT NULL, FOREIGN KEY(pid, steamid) REFERENCES gifts ON DELETE CASCADE, PRIMARY KEY(pid, steamid, tag));")
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetGiftTags(c context.Context, pid string, steamid int) ([]string, error) {
	var tags []string
	rows, err := dbpool.Query(c, "SELECT tag FROM tags WHERE pid = $1 AND steamid = $2;", pid, steamid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func CreateTags(c context.Context, gift Gift) error {
	for _, tag := range gift.Tags {
		_, err := dbpool.Exec(c, "INSERT INTO tags VALUES ($1, $2, $3);", gift.GifterID, gift.SteamID, tag)
		if err != nil {
			return err
		}
	}
	return nil
}
